package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
  Id        primitive.ObjectID `bson:"_id" mapper:"id"`
  Name      string             `bson:"name" mapper:"name"`
  CreatedAt time.Time          `bson:"createdAt" mapper:"createdAt"`
  UpdatedAt time.Time          `bson:"updatedAt" mapper:"updatedAt"`
}
