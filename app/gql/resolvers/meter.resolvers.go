package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *meterResolver) UpdatedBy(ctx context.Context, obj *models.Meter) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *meterResolver) CreatedBy(ctx context.Context, obj *models.Meter) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMeter(ctx context.Context, input models.NewMeter) (*models.Meter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meters(ctx context.Context, limit *int64, after *primitive.ObjectID) ([]*models.Meter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) CreateMeter(ctx context.Context, input models.NewMeter) (<-chan *models.Meter, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meter returns generated.MeterResolver implementation.
func (r *Resolver) Meter() generated.MeterResolver { return &meterResolver{r} }

type meterResolver struct{ *Resolver }
