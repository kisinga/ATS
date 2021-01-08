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
	Read(context.Context, string) (*models.Token, error)
	ReadByID(context.Context, primitive.ObjectID) (*models.Token, error)
	ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Token, error)
	Update(ctx context.Context, meterNumber string, newMeter models.Token) (*models.Token, error)
}

func NewRepository(database *storage.Database) Repository {
	return &repository{database}
}

type repository struct {
	db *storage.Database
}

func (r repository) Create(ctx context.Context, user models.Token) (*models.Token, error) {
	res, err := r.db.Client.Collection("tokens").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (r repository) Read(ctx context.Context, meterNumber string) (*models.Token, error) {
	user := models.Token{}
	return &user, r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&user)
}

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.Token, error) {
	user := models.Token{}
	return &user, r.db.Client.Collection("tokens").FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
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
func (r repository) Update(ctx context.Context, tokenNumber string, newToken models.Token) (*models.Token, error) {
	user := models.Token{}
	err := r.db.Client.Collection("tokens").FindOneAndUpdate(ctx, bson.M{"tokenNumber": tokenNumber}, newToken).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
