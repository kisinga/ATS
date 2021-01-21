package apiKey

import (
	"context"
	"errors"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kisinga/ATS/app/models"
)

var latestKey struct {
	mu  sync.RWMutex
	key *models.APIKey
}

type Interactor interface {
	GetLatest() *models.APIKey
	getLatestFromDB(ctx context.Context) (*models.APIKey, error)
	Generate(ctx context.Context, user models.User) (*models.APIKey, error)
	ListenForNew(ctx context.Context, consumer chan<- *models.APIKey)
}

type interactor struct {
	repository Repository
	listeners  map[primitive.ObjectID]chan<- *models.APIKey
	mu         sync.Mutex
}

func NewIterator(repo Repository) Interactor {
	kk := make(map[primitive.ObjectID]chan<- *models.APIKey, 0)
	i := &interactor{repository: repo, listeners: kk}

	// Fetch the latest key for in-memory cache
	go func() {
		key, err := i.getLatestFromDB(context.Background())
		if err != nil {
			log.Fatalln("Error fetching latest key for in-memory cache")
		}
		latestKey.mu.Lock()
		latestKey.key = key
		latestKey.mu.Unlock()
	}()
	go keyCreated(i, repo.apiKeyChan())
	return i

}

func keyCreated(i *interactor, channel chan *models.APIKey) {
	for {
		select {
		case key := <-channel:
			// Assign the new key to the private in-memory value
			latestKey.mu.Lock()
			latestKey.key = key
			latestKey.mu.Unlock()
			i.mu.Lock()
			for _, listener := range i.listeners {
				go func(l chan<- *models.APIKey) {
					l <- key
				}(listener)
			}
			i.mu.Unlock()
		}
	}
}

func (i *interactor) GetLatest() *models.APIKey {
	latestKey.mu.RLock()
	key := latestKey.key
	latestKey.mu.RUnlock()
	return key
}

func (i *interactor) ListenForNew(ctx context.Context, consumer chan<- *models.APIKey) {
	id := primitive.NewObjectID()
	i.mu.Lock()
	i.listeners[id] = consumer
	i.mu.Unlock()
	go func() {
		<-ctx.Done()
		i.mu.Lock()
		if _, ok := i.listeners[id]; ok {
			delete(i.listeners, id)
			i.mu.Unlock()
		}
	}()
}

func (i *interactor) getLatestFromDB(ctx context.Context) (*models.APIKey, error) {
	// create a limit of 1, then use the readmany func to get the latest value
	limit := int64(1)
	res, err := i.repository.ReadMany(ctx, primitive.NewObjectID(), &limit)
	if len(res) < 1 {
		return nil, errors.New("Nothing found")
	}
	return res[0], err
}

func (i *interactor) Generate(ctx context.Context, user models.User) (*models.APIKey, error) {
	key := models.APIKey{
		CreatedBy: user.ID,
		ID:        primitive.NewObjectID(),
	}
	return i.repository.Create(ctx, key)
}
