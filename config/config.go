package config

import (
	"context"
	"log"
	"os"
	"permission-api/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/joho/godotenv/autoload"
)

var MongoClient *mongo.Client

func ConnectDB() *mongo.Client {
	mongoUrl := os.Getenv("MONGO_URL")
	clientOpts := options.Client().ApplyURI(mongoUrl)
	var ctx = context.TODO()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	util.GreenLog("Connected to MongoDB!")

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
