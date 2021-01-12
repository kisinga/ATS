package app

import (
	"context"
	"fmt"
	"net/http"
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
	"github.com/kisinga/ATS/app/auth"
	"github.com/kisinga/ATS/app/gql/generated"
	"github.com/kisinga/ATS/app/gql/resolvers"
	"github.com/kisinga/ATS/app/registry"
	"github.com/kisinga/ATS/app/storage"
)

func Serve(db *storage.Database, firebase *firebase.App, port string, prod bool) error {
	// ctx := context.Background()
	gin := gin.Default()

	domain := registry.NewDomain(db)
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	apiRoutes := gin.Group("/api", auth.Mid(firebase), graphqlHandler(domain))
	{
		apiRoutes.POST("")
		apiRoutes.GET("")
	}
	gin.Use(auth.GinContextToContextMiddleware())

	gin.GET("/playground", playgroundHandler())

	return gin.Run(port)
}
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func graphqlHandler(domain *registry.Domain) gin.HandlerFunc {

	r := resolvers.NewResolver(domain)

	h := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			fmt.Println(initPayload)
			return ctx, nil
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
