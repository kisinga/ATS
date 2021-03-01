package token

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interactor interface {
	GetToken(ctx context.Context, ID primitive.ObjectID) (*models.Token, error)
	GetMany(ctx context.Context, meterNumber *string, after primitive.ObjectID, limit *int64, reversed bool) ([]*models.Token, error)
	AddToken(ctx context.Context, input models.NewToken, apiKey primitive.ObjectID) (*models.Token, error)
	UpdateTokenStatus(ctx context.Context, tokenID primitive.ObjectID, status models.TokenStatus) (*models.Token, error)
	// ListenForNew(ctx context.Context, consumer chan<- *models.Token)
	Count(ctx context.Context, query bson.M) (int64, error)
	AddListener(ctx context.Context, channel chan<- *models.Token, effectName TopicNames) primitive.ObjectID
}

type interactor struct {
	repository Repository
	effects    Effects
}

func NewIterator(repo Repository, effects Effects) Interactor {
	return &interactor{repo, effects}
}

func (i *interactor) AddListener(ctx context.Context, channel chan<- *models.Token, effectName TopicNames) primitive.ObjectID {
	return i.effects.Listeners().AddListener(ctx, channel, effectName)
}

func (i *interactor) Count(ctx context.Context, query bson.M) (int64, error) {
	return i.repository.Count(ctx, query)
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
