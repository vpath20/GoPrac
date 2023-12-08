package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	MongoDb := os.Getenv("MONGODB_URL")

	// =====Method1=====(Depriciated)

	// client ,err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// defer cancel()
	// fmt.Println("connection successfull")

	// =====Method2=====

	clientOptions := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}
