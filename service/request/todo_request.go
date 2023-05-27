package request

import "example/data/entity"

type TodoRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (req *TodoRequest) Map(data *entity.Todo) {
	if data == nil {
		panic("data is null")
	}

	data.Name = req.Name
}
