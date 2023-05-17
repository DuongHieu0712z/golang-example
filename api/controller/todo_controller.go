package controller

import (
	"example/common/errs"
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

		data := ctrl.usecase.GetPagedList(ctx, params)

		response.Response(ctx, http.StatusOK, data, nil)
	}
}

func (ctrl *todoController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get ID from param
		id := ctx.Param("id")

		data := ctrl.usecase.GetById(ctx, id)

		response.Response(ctx, http.StatusOK, data, nil)
	}
}

func (ctrl *todoController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Bind Todo form from body
		var form form.TodoForm
		if err := ctx.Bind(&form); err != nil {
			panic(errs.BadRequestError(err))
		}

		data := ctrl.usecase.Create(ctx, form)

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
			panic(errs.BadRequestError(err))
		}

		ctrl.usecase.Update(ctx, id, form)

		response.Response(ctx, http.StatusOK, "Success", nil)
	}
}

func (ctrl *todoController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get ID from param
		id := ctx.Param("id")

		ctrl.usecase.Delete(ctx, id)

		response.Response(ctx, http.StatusOK, "Success", nil)
	}
}
