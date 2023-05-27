package response

import (
	"example/data/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Id       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}

func ToUserResponse(data *entity.User) *UserResponse {
	if data == nil {
		panic("user is null")
	}

	return &UserResponse{
		Id:       data.Id,
		Username: data.Username,
	}
}
