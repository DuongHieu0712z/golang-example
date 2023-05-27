package base

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func (e *Entity) SetId(id interface{}) {
	e.Id = id.(primitive.ObjectID)
}

func (e *Entity) SetTime(isCreated bool) {
	now := time.Now()
	if isCreated {
		e.CreatedAt = now
	}
	e.UpdatedAt = now
}
