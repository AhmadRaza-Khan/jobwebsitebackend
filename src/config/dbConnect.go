package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	if err = client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("Could not ping MongoDB: %v", err))
	}

	fmt.Println("Connected to MongoDB successfully!")
	Client = client
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	if Client == nil {
		panic("MongoDB client is not initialized. Call ConnectDB first.")
	}
	return Client.Database("jobWebsite").Collection(collectionName)
}
