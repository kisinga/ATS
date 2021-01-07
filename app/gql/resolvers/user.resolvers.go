package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewMeter) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, limit *int64, after *string) ([]*models.User, error) {
	//Make sure that the provided limit doesnt exceed 50
	if *limit > 50 {
		*limit = int64(50)
	}
	//This step is very important, as fetching n+1 tuples always gives information about whether there is an element after the specified limit
	*limit = *limit + 1
	afterID := primitive.NewObjectID()
	if after != nil {
		var err error
		afterID, err = primitive.ObjectIDFromHex(*after)
		if err != nil {
			fmt.Println("Invalid users pagination id, ignoring")
		}
	}
	k, l := r.domain.User.GetMany(ctx, afterID, limit)
	return k, l
}
