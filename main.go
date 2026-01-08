package main

import (
	"go-test/db"
	"go-test/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
