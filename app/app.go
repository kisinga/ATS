package app

import (
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/kisinga/ATS/app/auth"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/gql/resolvers"
	"github.com/kisinga/ATS/app/registry"
	"github.com/kisinga/ATS/app/storage"
	"github.com/rs/cors"
)

func NewApp(db *storage.Database, firebase *firebase.App, port string, prod bool) error {
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4200",
			"ats-ke.web.app/",
		},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(auth.Middleware(firebase))

	domain := registry.NewDomain(db)

	r := resolvers.NewResolver(domain)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
	router.Handle("/playground", playground.Handler("GraphQL", "/api"))
	router.Handle("/api", srv)
	router.Handle("/sessionInit", auth.SessionInit(firebase))
	router.Handle("/sessionTerm", auth.SessionTerm())

	return http.ListenAndServe(port, router)
}
