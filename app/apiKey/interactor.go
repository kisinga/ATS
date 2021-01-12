package apiKey

import (
	"context"
	"errors"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kisinga/ATS/app/models"
)

type Interactor interface {
	GetLatest(ctx context.Context) (*models.APIKey, error)
	Generate(ctx context.Context, user models.User) (*models.APIKey, error)
	ListenForNew(ctx context.Context, consumer chan *models.APIKey)
}

type interactor struct {
	repository Repository
	listeners  map[primitive.ObjectID]chan *models.APIKey
	mu         sync.Mutex
}

func NewIterator(repo Repository) Interactor {
	kk := make(map[primitive.ObjectID]chan *models.APIKey, 0)
	i := &interactor{repository: repo, listeners: kk}
	go keyCreated(i, repo.apiKeyChan())
	return i

}

func keyCreated(i *interactor, channel chan *models.APIKey) {
	for {
		select {
		case key := <-channel:
			i.mu.Lock()
			for _, listener := range i.listeners {
				go func(l chan *models.APIKey) {
					l <- key
				}(listener)
			}
			i.mu.Unlock()
		}
	}
}

func (i *interactor) ListenForNew(ctx context.Context, consumer chan *models.APIKey) {
	id := primitive.NewObjectID()
	i.mu.Lock()
	i.listeners[id] = consumer
	i.mu.Unlock()

	// go func(id primitive.ObjectID) {
	// 	for {
	// 		select {
	// 		// case <-ctx.Done():
	// 		// 	if _, ok := i.listeners[id]; ok {
	// 		// 		i.mu.Lock()
	// 		// 		delete(i.listeners, id)
	// 		// 		i.mu.Unlock()
	// 		// 	}
	// 		// }
	// 	}
	// }(id)
}

func (i *interactor) GetLatest(ctx context.Context) (*models.APIKey, error) {
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
