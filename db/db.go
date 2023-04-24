package db

import (
	"context"
	"example/config"
	"example/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Database
}

func ConnectDb() (*Database, error) {
	mongoUri := config.GetMongoURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	name := config.GetDbName()
	db := &Database{db: client.Database(name)}

	// Create indexes is here
	db.createIndexes(ctx, model.CreateTodoIndexes())

	return db, nil
}

func (db *Database) GetCollection(name string) *mongo.Collection {
	return db.db.Collection(name)
}

func (db *Database) createIndexes(ctx context.Context, models ...mongo.IndexModel) {
	name, err := db.GetCollection("todo").Indexes().CreateMany(ctx, models)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Created indexes: %v\n", name)
}
