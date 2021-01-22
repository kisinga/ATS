package token

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(context.Context, models.Token) (*models.Token, error)
	// Read(context.Context, string) (*models.Token, error)
	ReadByID(context.Context, primitive.ObjectID) (*models.Token, error)
	ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Token, error)
	Update(ctx context.Context, newMeter models.Token) (*models.Token, error)
	tokenCreatedChan() chan *models.Token
}

func NewRepository(database *storage.Database) Repository {
	return &repository{database, make(chan *models.Token, 0)}
}

type repository struct {
	db           *storage.Database
	tokenCreated chan *models.Token
}

func (r repository) Create(ctx context.Context, token models.Token) (*models.Token, error) {
	_, err := r.db.Client.Collection("tokens").InsertOne(ctx, token)
	return &token, err
}
func (r repository) tokenCreatedChan() chan *models.Token {
	return r.tokenCreated
}

// func (r repository) Read(ctx context.Context, meterNumber string) (*models.Token, error) {
// 	token := models.Token{}
// 	return &token, r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&token)
// }

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.Token, error) {
	token := models.Token{}
	return &token, r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"_id": ID}).Decode(&token)
}

func (r repository) ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Token, error) {
	tokens := []*models.Token{}
	DataCursor, dataErr := r.db.Client.Collection("tokens").Find(ctx,
		bson.M{"_id": bson.M{"$lt": after}},
		&options.FindOptions{Limit: limit, Sort: bson.M{"_id": -1}})

	if dataErr != nil {
		return nil, dataErr
	}
	for DataCursor.Next(ctx) {
		elem := models.Token{}
		err := DataCursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, &elem)
	}
	return tokens, nil
}
func (r repository) Update(ctx context.Context, newToken models.Token) (*models.Token, error) {
	token := models.Token{}
	err := r.db.Client.Collection("tokens").FindOneAndUpdate(ctx, bson.M{"_id": newToken.ID}, newToken).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
