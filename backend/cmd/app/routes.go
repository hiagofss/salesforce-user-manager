package main

import (
	"net/http"
	"salesforce-user-manager/internal/infrastructure"

	"github.com/gin-gonic/gin"
)

// setupRoutes configures all application routes
func setupRoutes(router *gin.Engine, container *infrastructure.AppContainer) {
	// Health check endpoint - returns 200 if healthy, 503 if degraded
	router.GET("/health", func(c *gin.Context) {
		healthStatus := container.HealthChecker.Check(c.Request.Context())

		statusCode := http.StatusOK
		if !healthStatus.IsHealthy {
			statusCode = http.StatusServiceUnavailable
		}

		c.JSON(statusCode, healthStatus)
	})

	// API v1 routes
	api := router.Group("/api/v1")
	{
		setupUserRoutes(api, container)
		setupOrgRoutes(api, container)
	}
}

// setupUserRoutes configures user-related routes
func setupUserRoutes(api *gin.RouterGroup, container *infrastructure.AppContainer) {
	users := api.Group("/users")
	{
		users.GET("/:id", container.UserHandler.GetUser)
		users.GET("", container.UserHandler.GetAllUsers)
		users.POST("", container.UserHandler.CreateUser)
		users.PUT("/:id", container.UserHandler.UpdateUser)
		users.DELETE("/:id", container.UserHandler.DeleteUser)
	}
}

// setupOrgRoutes configures org-related routes
func setupOrgRoutes(api *gin.RouterGroup, container *infrastructure.AppContainer) {
	orgs := api.Group("/orgs")
	{
		orgs.GET("/:id", container.OrgHandler.GetOrg)
		orgs.GET("", container.OrgHandler.GetAllOrgs)
		orgs.POST("", container.OrgHandler.CreateOrg)
		orgs.PUT("/:id", container.OrgHandler.UpdateOrg)
		orgs.DELETE("/:id", container.OrgHandler.DeleteOrg)
		orgs.POST("/:id/sync-users", container.OrgHandler.SyncUsers)
	}
}
