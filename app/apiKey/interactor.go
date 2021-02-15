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
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	i := &interactor{repository: repo}
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

	// go UpdateCache(i, repo.apiKeyChan())

	return i

}

func UpdateCache(channel chan *models.APIKey) {
	for {
		select {
		case key := <-channel:
			// Assign the new key to the private in-memory value
			latestKey.mu.Lock()
			latestKey.key = key
			latestKey.mu.Unlock()
		}
	}
}

func (i *interactor) GetLatest() *models.APIKey {
	latestKey.mu.RLock()
	key := latestKey.key
	latestKey.mu.RUnlock()
	return key
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
