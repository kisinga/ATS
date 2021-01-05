package resolvers

import (
	"github.com/kisinga/ATS/app/meter"
	"github.com/kisinga/ATS/app/token"
	"github.com/kisinga/ATS/app/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	meter *meter.Interactor
	user  *user.Interactor
	token *token.Interactor
}

func NewResolver(meter *meter.Interactor, user *user.Interactor, token *token.Interactor) *Resolver {
	return &Resolver{meter, user, token}
}
