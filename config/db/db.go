package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	URI = "mongodb+srv://<username>:<password>@cicdcluster.4yy64.mongodb.net/<database>?retryWrites=true&w=majority"
)

func GetDBCollection() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// Get the collection
	collection := client.Database("LoginDatabase").Collection("Users")

	return collection, nil
}
