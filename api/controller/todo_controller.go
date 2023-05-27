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
	GetPagination() gin.HandlerFunc
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

func (ctrl todoController) GetPagination() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params pagination.PagingParams
		exchange.BindQuery(ctx, &params)

		resp := ctrl.usecase.GetPagination(ctx, params)

		exchange.ResponseSuccess(ctx, http.StatusOK, resp)
	}
}

func (ctrl todoController) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		resp := ctrl.usecase.GetById(ctx, id)

		exchange.ResponseSuccess(ctx, http.StatusOK, resp)
	}
}

func (ctrl todoController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.TodoRequest
		exchange.BindBody(ctx, &req)

		resp := ctrl.usecase.Create(ctx, req)

		exchange.ResponseSuccess(ctx, http.StatusCreated, resp)
	}
}

func (ctrl todoController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var req request.TodoRequest
		exchange.BindBody(ctx, &req)

		ctrl.usecase.Update(ctx, id, req)

		exchange.ResponseSuccess(ctx, http.StatusOK, nil)
	}
}

func (ctrl todoController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		ctrl.usecase.Delete(ctx, id)

		exchange.ResponseSuccess(ctx, http.StatusOK, nil)
	}
}
