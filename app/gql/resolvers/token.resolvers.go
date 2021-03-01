package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kisinga/ATS/app/domain/crudModels"
	"github.com/kisinga/ATS/app/domain/token"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *queryResolver) GetTokens(ctx context.Context, limit *int64, beforeOrAfter *primitive.ObjectID, reversed *bool, meterNumber *string) (*models.TokenConnection, error) {
	//Make sure that the provided limit doesnt exceed 50
	if *limit > 50 {
		*limit = int64(50)
	}
	//This step is very important, as fetching n+1 tuples always gives information about whether there is an element after the specified limit
	*limit = *limit + 1
	afterID := primitive.NewObjectID()
	if beforeOrAfter != nil {
		afterID = *beforeOrAfter
	}
	query := bson.M{}
	if meterNumber != nil {
		query = bson.M{
			"meterNumber": *meterNumber,
		}
	}
	count, err := r.domain.Token.Count(ctx, query)
	if err != nil {
		return nil, err
	}
	rev := false
	if reversed != nil {
		rev = *reversed
	}
	k, l := r.domain.Token.GetMany(ctx, meterNumber, afterID, limit, rev)
	connection := models.TokenConnection{
		Data: k,
	}
	connection.CreateConection(*limit, count)
	return &connection, l
}

func (r *subscriptionResolver) TokenCreated(ctx context.Context, meterNumber *string) (<-chan *models.Token, error) {
	channel := make(chan *models.Token)
	r.domain.Listeners.Token.AddListener(ctx, channel, token.TopicNames(crudModels.Create))
	return channel, nil
}

func (r *tokenResolver) Status(ctx context.Context, obj *models.Token) (int64, error) {
	return int64(int(obj.Status)), nil
}

// Token returns generated.TokenResolver implementation.
func (r *Resolver) Token() generated.TokenResolver { return &tokenResolver{r} }

type tokenResolver struct{ *Resolver }
