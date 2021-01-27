package token

import (
	"context"
	"sync"

	"github.com/kisinga/ATS/app/meter"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interactor interface {
	GetToken(ctx context.Context, ID primitive.ObjectID) (*models.Token, error)
	GetMany(ctx context.Context, meterNumber *string, after primitive.ObjectID, limit *int64, reversed bool) ([]*models.Token, error)
	AddToken(ctx context.Context, input models.NewToken, apiKey primitive.ObjectID) (*models.Token, error)
	UpdateTokenStatus(ctx context.Context, tokenID primitive.ObjectID, status models.TokenStatus) (*models.Token, error)
	ListenForNew(ctx context.Context, consumer chan<- *models.Token)
	Count(ctx context.Context, query bson.M) (int64, error)
}

type interactor struct {
	repository      Repository
	meterRepository meter.Repository
	listeners       map[primitive.ObjectID]chan<- *models.Token
	mu              sync.Mutex
}

func NewIterator(repo Repository, meterRepo meter.Repository) Interactor {
	kk := make(map[primitive.ObjectID]chan<- *models.Token, 0)
	i := &interactor{repository: repo, listeners: kk}
	go tokenCreated(i, repo.tokenCreatedChan())
	return i

}

func (i *interactor) Count(ctx context.Context, query bson.M) (int64, error) {
	return i.repository.Count(ctx, query)
}

func (i *interactor) ListenForNew(ctx context.Context, consumer chan<- *models.Token) {
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

func tokenCreated(i *interactor, channel chan *models.Token) {
	for {
		select {
		case key := <-channel:
			i.mu.Lock()
			for _, listener := range i.listeners {
				go func(l chan<- *models.Token) {
					l <- key
				}(listener)
			}
			i.mu.Unlock()
		}
	}
}
func (i *interactor) GetToken(ctx context.Context, ID primitive.ObjectID) (*models.Token, error) {
	return i.repository.ReadByID(ctx, ID)
}

func (i *interactor) AddToken(ctx context.Context, input models.NewToken, apiKey primitive.ObjectID) (*models.Token, error) {
	newToken := models.Token{
		ID:          primitive.NewObjectID(),
		MeterNumber: input.MeterNumber,
		APIKey:      apiKey,
		Status:      models.StatusNew,
		TokenString: input.TokenString,
	}

	return i.repository.Create(ctx, newToken)
}

func (i *interactor) UpdateTokenStatus(ctx context.Context, tokenID primitive.ObjectID, status models.TokenStatus) (*models.Token, error) {
	token, err := i.GetToken(ctx, tokenID)
	if err != nil {
		return nil, err
	}
	token.Status = status
	return i.repository.Update(ctx, *token)
}

func (i *interactor) GetMany(ctx context.Context, meterNumber *string, beforeOrAfter primitive.ObjectID, limit *int64, reversed bool) ([]*models.Token, error) {
	return i.repository.ReadMany(ctx, meterNumber, beforeOrAfter, limit, reversed)

}
