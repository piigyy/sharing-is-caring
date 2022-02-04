package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	client, clientErr := mongo.NewClient(options.Client().ApplyURI(uri))
	if clientErr != nil {
		return nil, clientErr
	}

	if connectErr := client.Connect(ctx); connectErr != nil {
		return nil, connectErr
	}

	return client, nil
}
