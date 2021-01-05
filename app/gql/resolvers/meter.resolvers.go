package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/models"
)

func (r *mutationResolver) CreateMeter(ctx context.Context, input models.NewMeter) (*models.Meter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meters(ctx context.Context, limit *int64, after *string) ([]*models.Meter, error) {
	panic(fmt.Errorf("not implemented"))
}
