package main

import (
	"time"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/my-game873206/my-game-data/graph"
	"gitlab.com/my-game873206/my-game-data/graph/generated"
	"gitlab.com/my-game873206/my-game-data/internal/db"
	"os"
	"strings"
)

func getAllowOrigins() []string {
    rawOrigins := os.Getenv("ALLOW_ORIGINS")
    
    if rawOrigins == "" {
        return []string{"http://localhost:3000"}
    }

    origins := strings.Split(rawOrigins, ",")

    for i := range origins {
        origins[i] = strings.TrimSpace(origins[i])
    }

    return origins
}


const defaultPort = "8080"

func main() {
	godotenv.Load(".env")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: getAllowOrigins(),
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Accept"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

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

	r.OPTIONS("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	r.GET("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
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
