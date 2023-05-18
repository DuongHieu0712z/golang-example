package response

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoResponse struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"createdAt"`
}
