package meter

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(context.Context, models.Meter) (*models.Meter, error)
	Read(context.Context, string) (*models.Meter, error)
	ReadByID(context.Context, primitive.ObjectID) (*models.Meter, error)
	ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Meter, error)
	Update(ctx context.Context, meterNumber string, newMeter models.Meter) (*models.Meter, error)
}

func NewRepository(database *storage.Database) Repository {
	return &repository{database}
}

type repository struct {
	db *storage.Database
}

func (r repository) Create(ctx context.Context, user models.Meter) (*models.Meter, error) {
	res, err := r.db.Client.Collection("meters").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return &user, nil
}

func (r repository) Read(ctx context.Context, meterNumber string) (*models.Meter, error) {
	user := models.Meter{}
	return &user, r.db.Client.Collection("meters").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&user)
}

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.Meter, error) {
	user := models.Meter{}
	return &user, r.db.Client.Collection("meters").FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
}

func (r repository) ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Meter, error) {
	meters := []*models.Meter{}
	DataCursor, dataErr := r.db.Client.Collection("meters").Find(ctx,
		bson.M{"_id": bson.M{"$lt": after}},
		&options.FindOptions{Limit: limit, Sort: bson.M{"_id": -1}})

	if dataErr != nil {
		return nil, dataErr
	}
	for DataCursor.Next(ctx) {
		elem := models.Meter{}
		err := DataCursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		meters = append(meters, &elem)
	}
	return meters, nil
}
func (r repository) Update(ctx context.Context, meterNumber string, newMeter models.Meter) (*models.Meter, error) {
	user := models.Meter{}
	err := r.db.Client.Collection("meters").FindOneAndUpdate(ctx, bson.M{"meterNumber": meterNumber}, newMeter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
