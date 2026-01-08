package routes

import (
	"go-test/internal/handlers"
	"go-test/internal/repository"
	"go-test/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	repo := repository.NewEventRepository()
	service := services.NewEventService(repo)
	handler := handlers.NewEventHandler(service)

	server.GET("/events", handler.GetEvents)
	server.GET("/events/:id", handler.GetEvent)
	server.POST("/events", handler.CreateEvent)
	server.PUT("/events/:id", handler.UpdateEvent)
	server.DELETE("/events/:id", handler.DeleteEvent)
}
