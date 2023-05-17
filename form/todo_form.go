package form

type TodoForm struct {
	Name string `json:"name" form:"name" binding:"required"`
}
