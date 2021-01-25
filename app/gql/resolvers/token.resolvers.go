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

func (r *queryResolver) GetTokens(ctx context.Context, limit *int64, after *primitive.ObjectID, meterNumber *string) (*models.TokenConnection, error) {
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
	k, l := r.domain.Token.GetMany(ctx, afterID, limit)
	connection := models.TokenConnection{
		Data: k,
	}
	connection.CreateConection(*limit, count)
	return &connection, l
}

func (r *subscriptionResolver) TokenCreated(ctx context.Context, meterNumber *string) (<-chan *models.Token, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tokenResolver) Status(ctx context.Context, obj *models.Token) (int64, error) {
	return int64(int(obj.Status)), nil
}

// Token returns generated.TokenResolver implementation.
func (r *Resolver) Token() generated.TokenResolver { return &tokenResolver{r} }

type tokenResolver struct{ *Resolver }
