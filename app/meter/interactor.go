package meter

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interactor interface {
	GetMeter(context.Context, string) (*models.Meter, error)
	GetMeterByID(context.Context, primitive.ObjectID) (*models.Meter, error)
	GetMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Meter, error)
	AddMeter(ctx context.Context, meter models.NewMeter, creatorID primitive.ObjectID) (*models.Meter, error)
	UpdateMeter(ctx context.Context, meterNumber string, newMeter models.Meter) (*models.Meter, error)
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{repo}
}

func (i *interactor) GetMeter(ctx context.Context, meterNumber string) (*models.Meter, error) {
	return i.repository.Read(ctx, meterNumber)
}

func (i *interactor) GetMeterByID(ctx context.Context, ID primitive.ObjectID) (*models.Meter, error) {
	return i.repository.ReadByID(ctx, ID)
}

func (i *interactor) GetMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.Meter, error) {
	return i.repository.ReadMany(ctx, after, limit)
}

func (i *interactor) AddMeter(ctx context.Context, meter models.NewMeter, creatorID primitive.ObjectID) (*models.Meter, error) {
	newMeter := models.Meter{
		CreatedBy:   creatorID,
		ID:          primitive.NewObjectID(),
		MeterNumber: meter.MeterNumber,
		Location:    meter.Location,
	}
	return i.repository.Create(ctx, newMeter)
}

func (i *interactor) UpdateMeter(ctx context.Context, meterNumber string, newMeter models.Meter) (*models.Meter, error) {
	return i.repository.Update(ctx, meterNumber, newMeter)
}
