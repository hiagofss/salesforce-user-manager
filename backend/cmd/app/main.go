package main

import (
	"salesforce-user-manager/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize all dependencies via container
	container := infrastructure.NewAppContainer()

	// Setup router
	router := gin.Default()

	// Configure all routes
	setupRoutes(router, container)

	router.Run() // listens on 0.0.0.0:8080 by default
}
