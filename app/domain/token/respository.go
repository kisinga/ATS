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
	ReadMany(ctx context.Context, meterNumber *string, after primitive.ObjectID, limit *int64, reversed bool) ([]*models.Token, error)
	Update(ctx context.Context, newMeter models.Token) (*models.Token, error)
	// tokenCreatedChan() chan *models.Token
	Count(ctx context.Context, query bson.M) (int64, error)
}

// CrudChannels keeps a list of channels to which values are emitted when certain CRUD operations are performed
type CrudChannels struct {
	Created chan models.Token
	Read    chan models.Token
	Updated chan models.Token
	Deleted chan models.Token
}

// NewCrudChannels creates an instance of CrudChannels
func NewCrudChannels() *CrudChannels {
	return &CrudChannels{
		Created: make(chan models.Token),
		Read:    make(chan models.Token),
		Updated: make(chan models.Token),
		Deleted: make(chan models.Token),
	}
}
func NewRepository(database *storage.Database, topics *Topics) Repository {
	return &repository{database, topics}
}

type repository struct {
	db     *storage.Database
	topics *Topics
}

func (r repository) Create(ctx context.Context, token models.Token) (*models.Token, error) {
	_, err := r.db.Client.Collection("tokens").InsertOne(ctx, token)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Created, token)
	return &token, err
}

// func (r repository) Read(ctx context.Context, meterNumber string) (*models.Token, error) {
// 	token := models.Token{}
// 	return &token, r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&token)
// }
func (r repository) Count(ctx context.Context, query bson.M) (int64, error) {
	return r.db.Client.Collection("tokens").CountDocuments(ctx, query)
}

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.Token, error) {
	token := models.Token{}
	err := r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"_id": ID}).Decode(&token)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Read, token)
	return &token, err
}

func (r repository) ReadMany(ctx context.Context, meterNumber *string, beforeOrAfter primitive.ObjectID, limit *int64, reversed bool) ([]*models.Token, error) {
	tokens := []*models.Token{}
	direcction := bson.M{"$lt": beforeOrAfter}
	sort := bson.M{"_id": -1}
	if reversed {
		direcction = bson.M{"$gte": beforeOrAfter}
		sort = bson.M{"_id": 1}
	}
	query := bson.M{"_id": direcction}
	if meterNumber != nil {
		query = bson.M{"_id": direcction, "meterNumber": *meterNumber}
	}
	DataCursor, dataErr := r.db.Client.Collection("tokens").Find(ctx,
		query,
		&options.FindOptions{Limit: limit, Sort: sort})

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
	// reverse the order so that pagination info is correct
	if reversed {
		for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
			tokens[i], tokens[j] = tokens[j], tokens[i]
		}
	}
	for _, token := range tokens {
		r.topics.Emit(r.topics.Read, *token)
	}
	return tokens, nil
}
func (r repository) Update(ctx context.Context, newToken models.Token) (*models.Token, error) {
	token := models.Token{}
	err := r.db.Client.Collection("tokens").FindOneAndUpdate(ctx, bson.M{"_id": newToken.ID}, newToken).Decode(&token)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Updated, token)
	return &token, nil
}
