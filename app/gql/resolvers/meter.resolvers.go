package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/handlers/auth"
	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *meterResolver) UpdatedBy(ctx context.Context, obj *models.Meter) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.UpdatedBy)
}

func (r *meterResolver) CreatedBy(ctx context.Context, obj *models.Meter) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.CreatedBy)
}

func (r *mutationResolver) CreateMeter(ctx context.Context, input models.NewMeter) (*models.Meter, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	return r.domain.Meter.AddMeter(ctx, input, me.ID)
}

func (r *mutationResolver) DisableMeter(ctx context.Context, meterNumber string) (*models.Meter, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	target, err := r.domain.Meter.GetMeter(ctx, meterNumber)
	if err != nil {
		return nil, err
	}
	target.Active = false
	target.UpdatedBy = me.ID
	return r.domain.Meter.UpdateMeter(ctx, meterNumber, *target)
}

func (r *mutationResolver) EnableMeter(ctx context.Context, meterNumber string) (*models.Meter, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	target, err := r.domain.Meter.GetMeter(ctx, meterNumber)
	if err != nil {
		return nil, err
	}
	target.Active = true
	target.UpdatedBy = me.ID
	return r.domain.Meter.UpdateMeter(ctx, meterNumber, *target)
}

func (r *queryResolver) Meters(ctx context.Context, limit *int64, after *primitive.ObjectID) (*models.MeterConnection, error) {
	//Make sure that the provided limit doesnt exceed 50
	if *limit > 50 {
		*limit = int64(50)
	}
	//This step is very important, as fetching n+1 tuples always gives information about whether there is an element after the specified limit
	*limit = *limit + 1
	afterID := primitive.NewObjectID()
	if after != nil {
		afterID = *after
	}
	k, l := r.domain.Meter.GetMany(ctx, afterID, limit)
	connection := models.MeterConnection{
		Data: k,
	}
	connection.CreateConection(*limit)
	return &connection, l
}

// Meter returns generated.MeterResolver implementation.
func (r *Resolver) Meter() generated.MeterResolver { return &meterResolver{r} }

type meterResolver struct{ *Resolver }
