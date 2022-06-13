package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const uri = "mongodb+srv://nattapong:018125337@cluster0.9tujx50.mongodb.net/test"

func ConnectMongo() *mongo.Client {
	if client != nil {
		return client
	}

	var err error

	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err.Error())
	}

	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err.Error())
	}

	return client
}
