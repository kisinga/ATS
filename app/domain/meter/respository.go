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
	Count(ctx context.Context) (int64, error)
}

// CrudChannels keeps a list of channels to which values are emitted when certain CRUD operations are performed
type CrudChannels struct {
	Created chan models.Meter
	Read    chan models.Meter
	Updated chan models.Meter
	Deleted chan models.Meter
}

// NewCrudChannels creates an instance of CrudChannels
func NewCrudChannels() *CrudChannels {
	return &CrudChannels{
		Created: make(chan models.Meter),
		Read:    make(chan models.Meter),
		Updated: make(chan models.Meter),
		Deleted: make(chan models.Meter),
	}
}
func NewRepository(database *storage.Database, topics *Topics) Repository {
	return &repository{database, topics}
}

type repository struct {
	db     *storage.Database
	topics *Topics
}

func (r repository) Create(ctx context.Context, meter models.Meter) (*models.Meter, error) {
	res, err := r.db.Client.Collection("meters").InsertOne(ctx, meter)
	if err != nil {
		return nil, err
	}
	meter.ID = res.InsertedID.(primitive.ObjectID)
	r.topics.Emit(r.topics.Created, meter)
	return &meter, nil
}

func (r repository) Count(ctx context.Context) (int64, error) {
	return r.db.Client.Collection("meters").CountDocuments(ctx, bson.M{})
}

func (r repository) Read(ctx context.Context, meterNumber string) (*models.Meter, error) {
	meter := models.Meter{}
	err := r.db.Client.Collection("meters").FindOne(ctx, bson.M{"meterNumber": meterNumber}).Decode(&meter)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Read, meter)
	return &meter, err
}

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.Meter, error) {
	meter := models.Meter{}
	err := r.db.Client.Collection("meters").FindOne(ctx, bson.M{"_id": ID}).Decode(&meter)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Read, meter)
	return &meter, err
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
	meter := models.Meter{}
	err := r.db.Client.Collection("meters").FindOneAndUpdate(ctx, bson.M{"meterNumber": meterNumber}, bson.M{"$set": newMeter}).Decode(&meter)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Updated, meter)
	return &meter, err
	return &meter, nil
}
