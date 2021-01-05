package app

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/gql/resolvers"
	"github.com/kisinga/ATS/app/storage"
)

func NewApp(d *storage.Database, port string, prod bool) error {
	router := chi.NewRouter()

	// router.Use(auth.Middleware(db))
	// meter := meter.NewIterator()
	r := resolvers.NewResolver(nil, nil, nil)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
	router.Handle("/playground", playground.Handler("GraphQL", "/api"))
	router.Handle("/api", srv)

	return http.ListenAndServe(port, router)
}
