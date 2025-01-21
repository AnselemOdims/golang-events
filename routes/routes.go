package routes

import (
	"event-planning/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("events", handleGetEvents)
	server.GET("events/:id", handleGetEventByID)

	authenticate := server.Group("/")

	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("events", handlePostEvents)
	authenticate.PUT("events/:id", handleUpdateEvent)
	authenticate.DELETE("events/:id", handleDeleteEvent)
	authenticate.GET("users", handleGetUsers)

	server.POST("signup", handleSignUp)
	server.POST("login", handleLogin)
}