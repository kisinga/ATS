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
}

// CrudChannels keeps a list of channels to which values are emitted when certain CRUD operations are performed
type CrudChannels struct {
	Created chan models.APIKey
	Read    chan models.APIKey
	Updated chan models.APIKey
	Deleted chan models.APIKey
}

// NewCrudChannels creates an instance of CrudChannels
func NewCrudChannels() *CrudChannels {
	return &CrudChannels{
		Created: make(chan models.APIKey),
		Read:    make(chan models.APIKey),
		Updated: make(chan models.APIKey),
		Deleted: make(chan models.APIKey),
	}
}

func NewRepository(database *storage.Database, topics *Topics) Repository {
	return &repository{database, topics}
}

type repository struct {
	db     *storage.Database
	topics *Topics
}

func (r repository) Create(ctx context.Context, apiKey models.APIKey) (*models.APIKey, error) {
	_, err := r.db.Client.Collection("apikeys").InsertOne(ctx, apiKey)
	// send to chanel via a goroutine so it doesnt block
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Created, apiKey)
	return &apiKey, err
}

// func (r repository) Read(ctx context.Context, meterNumber string) (*models.APIKey, error) {
// 	apiKey := models.APIKey{}
// 	return &apiKey, r.db.Client.Collection("apikeys").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&apiKey)
// }

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.APIKey, error) {
	apiKey := models.APIKey{}
	err := r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"_id": ID}).Decode(&apiKey)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Read, apiKey)
	return &apiKey, err
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
	for _, apiKey := range apikeys {
		r.topics.Emit(r.topics.Read, *apiKey)
	}
	return apikeys, nil
}
func (r repository) Update(ctx context.Context, newToken models.APIKey) (*models.APIKey, error) {
	apiKey := models.APIKey{}
	err := r.db.Client.Collection("apikeys").FindOneAndUpdate(ctx, bson.M{"_id": newToken.ID}, newToken).Decode(&apiKey)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Updated, apiKey)
	return &apiKey, nil
}
