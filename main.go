package main

import (
	"example/api/routes"
	"example/config"
	"example/db"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := casbin.NewEnforcer("./rbac/model.conf", "./rbac/policy.csv")
	if err != nil {
		log.Fatalln(err)
	}

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

	host := config.GetHost()
	log.Println("Connect database successfully...")
	log.Printf("Start server at http://%s", host)

	if err := router.Run(host); err != nil {
		log.Fatalln(err)
	}
}
