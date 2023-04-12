package controller

import (
	"example/common/response"
	"example/db"
	"example/form"
	"example/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type todoController struct {
	usecase usecase.TodoUsecase
}

func NewTodoController(db *db.Database) TodoController {
	return &todoController{
		usecase: usecase.NewTodoUsecase(db),
	}
}

func (ctrl *todoController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := ctrl.usecase.GetById(ctx, id)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, nil, err)
		return
	}

	response.Response(ctx, http.StatusOK, data, nil)
}

func (ctrl *todoController) Create(ctx *gin.Context) {
	var form form.TodoForm
	if err := ctx.BindJSON(&form); err != nil {
		response.Response(ctx, http.StatusBadRequest, nil, err)
		return
	}

	data, err := ctrl.usecase.Create(ctx, form)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, nil, err)
		return
	}

	response.Response(ctx, http.StatusCreated, data, nil)
}

func (ctrl *todoController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var form form.TodoForm
	if err := ctx.BindJSON(&form); err != nil {
		response.Response(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if err := ctrl.usecase.Update(ctx, id, form); err != nil {
		response.Response(ctx, http.StatusBadRequest, nil, err)
		return
	}

	response.Response(ctx, http.StatusOK, "Success", nil)
}

func (ctrl *todoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := ctrl.usecase.Delete(ctx, id); err != nil {
		response.Response(ctx, http.StatusBadRequest, nil, err)
		return
	}

	response.Response(ctx, http.StatusOK, "Success", nil)
}
