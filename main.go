package main

import (
	"example/api/middleware"
	"example/api/routes"
	"example/config"
	"example/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middleware.Recovery())

	db, err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	group := router.Group("/api/v1")

	routes.CreateAuthRoutes(group, db)
	routes.CreateTodoRoutes(group, db)

	port := config.GetPort()
	log.Println("Connect database successfully...")
	log.Printf("Start server at http://localhost:%s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalln(err)
	}
}
