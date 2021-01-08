package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/models"
)

func (r *aPIKeyResolver) UpdatedBy(ctx context.Context, obj *models.APIKey) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *aPIKeyResolver) CreatedBy(ctx context.Context, obj *models.APIKey) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Generate(ctx context.Context) (*models.APIKey, error) {
	panic(fmt.Errorf("not implemented"))
}

// APIKey returns generated.APIKeyResolver implementation.
func (r *Resolver) APIKey() generated.APIKeyResolver { return &aPIKeyResolver{r} }

type aPIKeyResolver struct{ *Resolver }
