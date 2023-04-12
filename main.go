package main

import (
	"example/api/routes"
	"example/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	db, err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	router := r.Group("/api/v1")

	routes.CreateTodoRoutes(router, db)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
