package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func ConnectDB() *mongo.Client {
	uri:=os.Getenv("DB_URI")
	if uri == ""{
		uri="mongodb://localhost:27017"
	}

	ctx, cancel:=context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err :=mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err!=nil{
		log.Fatal("Database connection error", err)
	}

	err = client.Ping(ctx, nil)
	if err!=nil{
		log.Fatal("Database connection error")
	}
	log.Println("Database connection successful")
	return client
}