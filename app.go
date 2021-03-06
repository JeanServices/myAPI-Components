package jeanservices

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MyMongo(atlasURI string) (*mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI(atlasURI))
	if err != nil {
		panic(err.Error())
	}

	err = client.Connect(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}

	return client
}
