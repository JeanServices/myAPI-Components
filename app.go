package jeanservices

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID		string	`bson:"id"`
	Name	string	`bson:"name"`
}

func MyMongo(atlasURI string) (*mongo.Client) {

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

func InsertUser(client *mongo.Client, id string, name string) {

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
