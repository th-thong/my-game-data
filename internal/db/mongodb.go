package db

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectDB() *mongo.Client {
    uri := os.Getenv("DB_URI")
    if uri == "" {
        uri = "mongodb://localhost:27017"
    }

    client, err := mongo.Connect(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Database connection initialization error: ", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Database connection failed (Ping): ", err)
    }

    log.Println("Database connection successful")
    return client
}