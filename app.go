package jeanservices

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MyMongo() (*mongo.Client) {
	atlasURI := "mongodb+srv://myapi:pri@jean.zjitk.mongodb.net/myAPI?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(atlasURI))
	if err != nil {
		panic(err.Error())
	}

	err = client.Connect(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	return client
}

func insertUser(client *mongo.Client, id string, name string) {

	_user := User{
		ID: id,
		Name: name,
	}

	_, err := client.Database("myAPI").Collection("users").InsertOne(context.TODO(), _user)
	if err != nil {
		panic(err)
	}

	return
}
