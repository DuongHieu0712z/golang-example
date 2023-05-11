package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
