package app

import (
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	domain := registry.NewDomain(db)

	r := resolvers.NewResolver(domain)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
	router.Handle("/playground", playground.Handler("GraphQL", "/api"))
	router.Route("/api", func(r chi.Router) {
		router.Use(auth.Middleware(firebase))
		router.Handle("/api", srv)
	})
	router.Post("/sessionInit", auth.SessionInit(firebase))
	router.Get("/sessionTerm", auth.SessionTerm())

	return http.ListenAndServe(port, router)
}
