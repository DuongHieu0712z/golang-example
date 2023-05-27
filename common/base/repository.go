package base

import (
	"example/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	name       string
	Db         *db.Database
	Collection *mongo.Collection
}

func NewRepository(db *db.Database, name string) *Repository {
	return &Repository{
		name:       name,
		Db:         db,
		Collection: db.GetCollection(name),
	}
}
