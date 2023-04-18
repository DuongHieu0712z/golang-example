package main

import (
	"example/api/routes"
	"example/config"
	"example/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Create log file
	file := config.CreateLogFile()
	router.Use(gin.LoggerWithWriter(file))

	db, err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	group := router.Group("/api/v1")

	routes.CreateTodoRoutes(group, db)

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}
}
