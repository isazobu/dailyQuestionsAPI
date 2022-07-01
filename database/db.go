package db

import (
	"fmt"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	fmt.Println("ConnectDB the GetCollection process")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://faruk:faruk@cluster0.hbxro.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println("Connection Error 1 ")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Connection Error 1 ")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Connection error")
	}

	fmt.Println("Connected to MongoDB")
	return client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("FIXME").Collection(collectionName)
	fmt.Println("GetCollection successful")
	return collection
}
