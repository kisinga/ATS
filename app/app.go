package app

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	gqlLru "github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kisinga/ATS/app/domain"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/gql/resolvers"
	"github.com/kisinga/ATS/app/handlers/auth"
	handlers "github.com/kisinga/ATS/app/handlers/token"
	"github.com/kisinga/ATS/app/storage"
)

func Serve(db *storage.Database, firebase *firebase.App, port string, prod bool) error {
	// ctx := context.Background()
	router := gin.Default()
	router.Use(gin.Recovery()) // add Recovery middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(auth.GinContextToContextMiddleware())

	domain := domain.New(db)
	apiRoutes := router.Group("/api",
		auth.Middleware(firebase),
		graphqlHandler(domain, firebase))
	{
		apiRoutes.Any("")
	}
	router.GET("/playground", playgroundHandler())
	router.POST("/sessionInit", auth.SessionInit(firebase, domain))
	router.POST("/token", handlers.TokenHandler(domain))
	router.GET("/sessionTerm", auth.SessionTerm())
	// listenForNewTokens(domain)
	return router.Run(port)
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func graphqlHandler(domain *domain.Domain, fb *firebase.App) gin.HandlerFunc {

	r := resolvers.NewResolver(domain)

	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, payload transport.InitPayload) (context.Context, error) {
			if _, ok := payload["Authorization"].(string); !ok {
				return ctx, errors.New("No auth in payload")
			}
			bearer := payload["Authorization"].(string)
			token := ""
			if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
				token = bearer[7:]
			}
			user, err := auth.GetUserFromToken(ctx, fb, token)
			if err != nil {
				return ctx, err
			}
			ginContext, err := auth.GinContextFromContext(ctx)
			if err != nil {
				return ctx, err
			}
			ginContext.Set("user", user)
			return auth.GinContextToContext(ginContext), nil
		},
	})

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.SetQueryCache(gqlLru.New(1000))
	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: gqlLru.New(100),
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
