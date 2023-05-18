package controller

import (
	"example/common/exchange"
	"example/common/pagination"
	"example/db"
	"example/service/request"
	"example/service/usecase"
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
		params := pagination.GetPagingParams(ctx)

		data := ctrl.usecase.GetPagedList(ctx, params)

		exchange.ResponseSuccess(ctx, http.StatusOK, data)
	}
}

func (ctrl *todoController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		data := ctrl.usecase.GetById(ctx, id)

		exchange.ResponseSuccess(ctx, http.StatusOK, data)
	}
}

func (ctrl *todoController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request request.TodoRequest
		exchange.Bind(ctx, &request)

		data := ctrl.usecase.Create(ctx, request)

		exchange.ResponseSuccess(ctx, http.StatusCreated, data)
	}
}

func (ctrl *todoController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var request request.TodoRequest
		exchange.Bind(ctx, &request)

		ctrl.usecase.Update(ctx, id, request)

		exchange.ResponseSuccess(ctx, http.StatusOK, nil)
	}
}

func (ctrl *todoController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		ctrl.usecase.Delete(ctx, id)

		exchange.ResponseSuccess(ctx, http.StatusOK, nil)
	}
}
