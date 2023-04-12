package db

import (
	"context"
	"example/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Database
}

func ConnectDb() (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetMongoURI()))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return &Database{db: client.Database("test")}, nil
}

func (db *Database) GetCollection(name string) *mongo.Collection {
	return db.db.Collection(name)
}
