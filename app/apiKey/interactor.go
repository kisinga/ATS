package apiKey

import (
	"context"

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
}

func NewIterator(repo Repository) Interactor {
	kk := make(map[primitive.ObjectID]chan *models.APIKey, 0)
	go keyCreated(kk, repo.apiKeyChan())
	return &interactor{repo, kk}

}

func keyCreated(listeners map[primitive.ObjectID]chan *models.APIKey, channel chan *models.APIKey) {
	for {
		select {
		case key := <-channel:
			for _, listener := range listeners {
				go func(l chan *models.APIKey) {
					l <- key
				}(listener)
			}
		}
	}
}

func (i *interactor) ListenForNew(ctx context.Context, consumer chan *models.APIKey) {
	id := primitive.NewObjectID()
	i.listeners[id] = consumer
	go func(id primitive.ObjectID) {
		for {
			select {
			case <-ctx.Done():
				if _, ok := i.listeners[id]; ok {
					delete(i.listeners, id)
				}
			}
		}
	}(id)
}

func (i *interactor) GetLatest(ctx context.Context) (*models.APIKey, error) {
	// create a limit of 1, then use the readmany func to get the latest value
	limit := int64(1)
	res, err := i.repository.ReadMany(ctx, primitive.NewObjectID(), &limit)
	return res[0], err
}

func (i *interactor) Generate(ctx context.Context, user models.User) (*models.APIKey, error) {
	key := models.APIKey{
		CreatedBy: user.ID,
		ID:        primitive.NewObjectID(),
	}
	return i.repository.Create(ctx, key)
}
