package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Setup router
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
