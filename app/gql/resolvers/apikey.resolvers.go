package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kisinga/ATS/app/domain/apiKey"
	"github.com/kisinga/ATS/app/domain/crudModels"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/handlers/auth"
	"github.com/kisinga/ATS/app/models"
)

func (r *aPIKeyResolver) CreatedBy(ctx context.Context, obj *models.APIKey) (*models.User, error) {
	return r.domain.User.GetUserByID(ctx, obj.CreatedBy)
}

func (r *mutationResolver) GenerateAPIKey(ctx context.Context) (*models.APIKey, error) {
	me, err := auth.GetUserFromContext(ctx, r.domain)
	if err != nil {
		return nil, err
	}
	return r.domain.APIKey.Generate(ctx, *me)
}

func (r *queryResolver) CurrentAPIKey(ctx context.Context) (*models.APIKey, error) {
	return r.domain.APIKey.GetLatest(), nil
}

func (r *subscriptionResolver) APIKeyChanged(ctx context.Context) (<-chan *models.APIKey, error) {
	channel := make(chan *models.APIKey)
	r.domain.APIKey.AddListener(ctx, channel, apiKey.TopicNames(crudModels.Create))
	return channel, nil
}

// APIKey returns generated.APIKeyResolver implementation.
func (r *Resolver) APIKey() generated.APIKeyResolver { return &aPIKeyResolver{r} }

type aPIKeyResolver struct{ *Resolver }
