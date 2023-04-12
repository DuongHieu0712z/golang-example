package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoDto struct {
	Id   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}
