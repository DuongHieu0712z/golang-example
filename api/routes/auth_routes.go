package routes

import (
	"example/api/controller"
	"example/db"

	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(router *gin.RouterGroup, db *db.Database) {
	ctrl := controller.NewAuthController(db)

	group := router.Group("/auth")
	{
		group.POST("/register", ctrl.Register())
	}
}
