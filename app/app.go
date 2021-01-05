package app

import (
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/gql/resolvers"
	"github.com/kisinga/ATS/app/registry"
	"github.com/kisinga/ATS/app/storage"
)

func NewApp(db *storage.Database, firebase *firebase.App, port string, prod bool) error {
	router := chi.NewRouter()
	domain := registry.NewDomain(db, firebase)
	r := resolvers.NewResolver(domain)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
	router.Handle("/playground", playground.Handler("GraphQL", "/api"))
	router.Handle("/api", srv)

	return http.ListenAndServe(port, router)
}
