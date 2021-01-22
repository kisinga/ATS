package apiKey

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(context.Context, models.APIKey) (*models.APIKey, error)
	// Read(context.Context, string) (*models.APIKey, error)
	ReadByID(context.Context, primitive.ObjectID) (*models.APIKey, error)
	ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.APIKey, error)
	Update(ctx context.Context, newMeter models.APIKey) (*models.APIKey, error)
	apiKeyChan() chan *models.APIKey
}

func NewRepository(database *storage.Database) Repository {
	return &repository{database, make(chan *models.APIKey, 0)}
}

type repository struct {
	db            *storage.Database
	apiKeyCreated chan *models.APIKey
}

func (r repository) Create(ctx context.Context, token models.APIKey) (*models.APIKey, error) {
	_, err := r.db.Client.Collection("apikeys").InsertOne(ctx, token)
	// send to chanel via a goroutine so it doesnt block
	if err == nil {
		go func() {
			r.apiKeyCreated <- &token
		}()
	}
	return &token, err
}

func (r repository) apiKeyChan() chan *models.APIKey {
	return r.apiKeyCreated
}

// func (r repository) Read(ctx context.Context, meterNumber string) (*models.APIKey, error) {
// 	token := models.APIKey{}
// 	return &token, r.db.Client.Collection("apikeys").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&token)
// }

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.APIKey, error) {
	token := models.APIKey{}
	return &token, r.db.Client.Collection("apikeys").FindOne(ctx, bson.M{"_id": ID}).Decode(&token)
}

func (r repository) ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.APIKey, error) {
	apikeys := []*models.APIKey{}
	DataCursor, dataErr := r.db.Client.Collection("apikeys").Find(ctx,
		bson.M{"_id": bson.M{"$lt": after}},
		&options.FindOptions{Limit: limit, Sort: bson.M{"_id": -1}})

	if dataErr != nil {
		return nil, dataErr
	}
	for DataCursor.Next(ctx) {
		elem := models.APIKey{}
		err := DataCursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		apikeys = append(apikeys, &elem)
	}
	return apikeys, nil
}
func (r repository) Update(ctx context.Context, newToken models.APIKey) (*models.APIKey, error) {
	token := models.APIKey{}
	err := r.db.Client.Collection("apikeys").FindOneAndUpdate(ctx, bson.M{"_id": newToken.ID}, newToken).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
