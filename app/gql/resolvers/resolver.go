//go:generate go run github.com/99designs/gqlgen

package resolvers

import (
	"github.com/kisinga/ATS/app/registry"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	domain *registry.Domain
}

func NewResolver(domain *registry.Domain) *Resolver {
	return &Resolver{domain}
}
