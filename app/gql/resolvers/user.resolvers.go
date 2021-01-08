package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	me := ForContext(ctx)
	if me == nil {
		return nil, errors.New("failed extracting user from context")
	}
	return r.domain.User.AddUser(ctx, input, me.Email)
}

func (r *mutationResolver) DisableUser(ctx context.Context, email string) (*models.User, error) {
	user, err := r.domain.User.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	user.Disabled = true
	return r.domain.User.UpdateUser(ctx, email, *user)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, limit *int64, after *primitive.ObjectID) ([]*models.User, error) {
	//Make sure that the provided limit doesnt exceed 50
	if *limit > 50 {
		*limit = int64(50)
	}
	//This step is very important, as fetching n+1 tuples always gives information about whether there is an element after the specified limit
	*limit = *limit + 1
	afterID := primitive.NewObjectID()
	if after != nil {
		if after != nil {
			afterID = *after
		}
	}
	k, l := r.domain.User.GetMany(ctx, afterID, limit)
	return k, l
}

func (r *userResolver) UpdatedBy(ctx context.Context, obj *models.User) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.UpdatedBy)
}

func (r *userResolver) CreatedBy(ctx context.Context, obj *models.User) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.CreatedBy)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
