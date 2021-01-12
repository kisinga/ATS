package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/auth"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/models"
)

func (r *aPIKeyResolver) CreatedBy(ctx context.Context, obj *models.APIKey) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.CreatedBy)
}

func (r *mutationResolver) Generate(ctx context.Context) (*models.APIKey, error) {
	me, err := auth.GetUserIDFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	return r.domain.APIKey.Generate(ctx, *me)
}

func (r *subscriptionResolver) APIKeyChanged(ctx context.Context) (<-chan *models.APIKey, error) {
	panic(fmt.Errorf("not implemented"))
}

// APIKey returns generated.APIKeyResolver implementation.
func (r *Resolver) APIKey() generated.APIKeyResolver { return &aPIKeyResolver{r} }

type aPIKeyResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *subscriptionResolver) KeyCreated(ctx context.Context) (<-chan *models.APIKey, error) {
	kk := make(chan *models.APIKey)
	r.domain.APIKey.ListenForNew(ctx, kk)
	return kk, nil
}
