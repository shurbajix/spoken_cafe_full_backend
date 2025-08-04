package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDataBase(uri string) *mongo.Client {
	if uri == "" {
		log.Println("No database URL provided, skipping database connection")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ClientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, ClientOptions)

	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return nil
	}

	log.Println("Connected to MongoDB")
	return client
}
