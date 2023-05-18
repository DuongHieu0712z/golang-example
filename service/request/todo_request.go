package request

type TodoRequest struct {
	Name string `json:"name,omitempty" form:"name,omitempty" binding:"required"`
}
