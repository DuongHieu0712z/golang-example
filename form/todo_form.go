package form

type TodoForm struct {
	Name string `json:"name" mapper:"name"`
}
