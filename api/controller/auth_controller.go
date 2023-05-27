package controller

import (
	"example/common/exchange"
	"example/db"
	"example/service/request"
	"example/service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register() gin.HandlerFunc
}

type authController struct {
	usecase usecase.AuthUsecase
}

func NewAuthController(db *db.Database) AuthController {
	return &authController{
		usecase: usecase.NewAuthUsecase(db),
	}
}

func (ctrl authController) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.RegisterRequest
		exchange.BindBody(ctx, &req)

		resp := ctrl.usecase.Register(ctx, req)

		exchange.ResponseSuccess(ctx, http.StatusOK, resp)
	}
}
