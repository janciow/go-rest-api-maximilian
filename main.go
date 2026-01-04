package main

import (
	"go-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetEvents()
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {

}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = "123"
	event.UserID = "user_1"
	event.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "Event created"})

}
