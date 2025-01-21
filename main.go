package main

import (
	"event-planning/db"
	"event-planning/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	db.InitDb()
	server.Run(":8000")
}

