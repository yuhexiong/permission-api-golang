package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/joho/godotenv/autoload"
)

var MongoClient *mongo.Client

func ConnectDB() *mongo.Client {
	mongoUrl := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(mongoUrl)
	var ctx = context.TODO()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	MongoClient = client
	return client
}

// Client instance
func GetDB() *mongo.Client {
	if MongoClient == nil {
		MongoClient = ConnectDB()
	}
	return MongoClient
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
