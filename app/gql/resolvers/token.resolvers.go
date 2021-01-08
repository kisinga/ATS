package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *queryResolver) Tokens(ctx context.Context, limit *int64, after *primitive.ObjectID) (*models.TokenConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) TokenCreated(ctx context.Context, meterNumber *string) (<-chan *models.Token, error) {
	panic(fmt.Errorf("not implemented"))
}
