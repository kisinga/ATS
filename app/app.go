package app

import (
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/kisinga/ATS/app/auth"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/gql/resolvers"
	"github.com/kisinga/ATS/app/registry"
	"github.com/kisinga/ATS/app/storage"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func NewApp(db *storage.Database, firebase *firebase.App, port string, prod bool) error {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	router := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{
			"http://localhost:4200",
			"https://ats-ke.web.app",
		},
		AllowCredentials: true,
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
		Debug:          !prod,
	}))

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))

	domain := registry.NewDomain(db)

	r := resolvers.NewResolver(domain)
	gqlHandler := generated.NewExecutableSchema(generated.Config{Resolvers: r})
	srv := handler.NewDefaultServer(gqlHandler)
	router.Handle("/playground", playground.Handler("GraphQL", "/api"))
	router.Group(func(rr chi.Router) {
		rr.Use(auth.Middleware(firebase))
		rr.Handle("/api", srv)
	},
	)
	router.Post("/sessionInit", auth.SessionInit(firebase, domain))
	router.Get("/sessionTerm", auth.SessionTerm())

	return http.ListenAndServe(port, router)
}
