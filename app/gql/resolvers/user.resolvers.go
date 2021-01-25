package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/handlers/auth"
	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, errors.New("failed extracting user from context")
	}
	return r.domain.User.AddUser(ctx, input, me.Email)
}

func (r *mutationResolver) DisableUser(ctx context.Context, email string) (*models.User, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	target, err := r.domain.User.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	target.Active = false
	target.UpdatedBy = me.ID
	return r.domain.User.UpdateUser(ctx, email, *target)
}

func (r *mutationResolver) EnableUser(ctx context.Context, email string) (*models.User, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	target, err := r.domain.User.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	target.Active = true
	target.UpdatedBy = me.ID
	return r.domain.User.UpdateUser(ctx, email, *target)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	target, err := r.domain.User.GetUser(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	target.Email = input.Email
	target.Name = input.Name
	target.UpdatedBy = me.ID
	return r.domain.User.UpdateUser(ctx, input.Email, *target)
}

func (r *queryResolver) Users(ctx context.Context, limit *int64, after *primitive.ObjectID) (*models.UsersConnection, error) {
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
	count, err := r.domain.Token.Count(ctx)
	if err != nil {
		return nil, err
	}
	k, l := r.domain.User.GetMany(ctx, afterID, limit)
	connection := models.UsersConnection{
		Data: k,
	}
	connection.CreateConection(*limit, count)
	return &connection, l
}

func (r *userResolver) UpdatedBy(ctx context.Context, obj *models.User) (*models.User, error) {
	if obj.UpdatedBy.IsZero() {
		return models.EmptyUser(), nil
	}
	return r.domain.User.GetUserByID(ctx, obj.UpdatedBy)
}

func (r *userResolver) CreatedBy(ctx context.Context, obj *models.User) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.CreatedBy)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
