package controller

import (
	"example/common/pagination"
	"example/common/response"
	"example/db"
	"example/form"
	"example/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	GetPagedList() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type todoController struct {
	usecase usecase.TodoUsecase
}

func NewTodoController(db *db.Database) TodoController {
	return &todoController{
		usecase: usecase.NewTodoUsecase(db),
	}
}

func (ctrl *todoController) GetPagedList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get paging params from query
		params := pagination.GetPagingParams(ctx)

		data, err := ctrl.usecase.GetPagedList(ctx, params)
		if err != nil {
			response.Response(ctx, http.StatusBadRequest, nil, err)
			return
		}

		response.Response(ctx, http.StatusOK, data, nil)
	}
}

func (ctrl *todoController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get ID from param
		id := ctx.Param("id")

		data, err := ctrl.usecase.GetById(ctx, id)
		if err != nil {
			response.Response(ctx, http.StatusBadRequest, nil, err)
			return
		}

		response.Response(ctx, http.StatusOK, data, nil)
	}
}

func (ctrl *todoController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Todo form from body
		var form form.TodoForm
		if err := ctx.Bind(&form); err != nil {
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
}

func (ctrl *todoController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get ID from param
		id := ctx.Param("id")
		// Bind Todo form from body
		var form form.TodoForm
		if err := ctx.Bind(&form); err != nil {
			response.Response(ctx, http.StatusBadRequest, nil, err)
			return
		}

		if err := ctrl.usecase.Update(ctx, id, form); err != nil {
			response.Response(ctx, http.StatusBadRequest, nil, err)
			return
		}

		response.Response(ctx, http.StatusOK, "Success", nil)
	}
}

func (ctrl *todoController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get ID from param
		id := ctx.Param("id")

		if err := ctrl.usecase.Delete(ctx, id); err != nil {
			response.Response(ctx, http.StatusBadRequest, nil, err)
			return
		}

		response.Response(ctx, http.StatusOK, "Success", nil)
	}
}
