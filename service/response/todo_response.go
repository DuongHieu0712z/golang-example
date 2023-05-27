package response

import (
	"example/data/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoResponse struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	CreatedAt time.Time          `json:"createdAt"`
}

func ToTodoResponse(data *entity.Todo) *TodoResponse {
	if data == nil {
		panic("data is null")
	}

	return &TodoResponse{
		Id:        data.Id,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}

func ToTodoResponseSlice(data []entity.Todo) []TodoResponse {
	res := []TodoResponse{}
	for _, val := range data {
		item := ToTodoResponse(&val)
		res = append(res, *item)
	}
	return res
}
