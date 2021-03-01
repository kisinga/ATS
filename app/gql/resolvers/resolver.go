//go:generate go run github.com/99designs/gqlgen

package resolvers

import (
	"github.com/kisinga/ATS/app/registry"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
type Resolver struct {
	domain *registry.Domain
}

func NewResolver(domain *registry.Domain) *Resolver {
	return &Resolver{domain}
}
