package routes

import (
	"example/api/controller"
	"example/db"

	"github.com/gin-gonic/gin"
)

func CreateTodoRoutes(router *gin.RouterGroup, db *db.Database) {
	ctrl := controller.NewTodoController(db)

	group := router.Group("/todos")
	{
		group.GET("", ctrl.GetPagination())
		group.POST("", ctrl.Create())
		group.GET("/:id", ctrl.GetById())
		group.PUT("/:id", ctrl.Update())
		group.DELETE("/:id", ctrl.Delete())
	}
}
