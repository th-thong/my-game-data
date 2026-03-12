package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/my-game873206/my-game-data/graph"
	"gitlab.com/my-game873206/my-game-data/graph/generated"
	"gitlab.com/my-game873206/my-game-data/internal/db"
	"gitlab.com/my-game873206/my-game-data/middleware"
	"os"
)


const defaultPort = "8080"

func main() {
	godotenv.Load(".env")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	mongoClient := db.ConnectDB()
	database := mongoClient.Database(os.Getenv("DB_NAME"))

	srv := handler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{DB: database},
			},
		),
	)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.Use(extension.Introspection{})

	r.POST("/query", middleware.AuthMiddleware(), func(c *gin.Context) {
		uid, _ := c.Get("user_id")
		ctx := context.WithValue(c.Request.Context(), "user_id", uid)
		srv.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	})

	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r.Run(":" + port)
}
