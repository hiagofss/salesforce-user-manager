package infrastructure

import (
	"salesforce-user-manager/internal/handler"
	"salesforce-user-manager/internal/repository"
	"salesforce-user-manager/internal/usecase"
)

// AppContainer holds all application dependencies
type AppContainer struct {
	// Repositories
	UserRepository repository.UserRepository
	OrgRepository  repository.OrgRepository

	// Usecases
	UserUsecase *usecase.UserUsecase
	OrgUsecase  *usecase.OrgUsecase

	// Handlers
	UserHandler *handler.UserHandler
	OrgHandler  *handler.OrgHandler

	// Health
	HealthChecker *HealthChecker
}

// NewAppContainer initializes all dependencies
func NewAppContainer() *AppContainer {
	c := &AppContainer{}
	c.initializeUser()
	c.initializeOrg()
	c.HealthChecker = NewHealthChecker(c)
	return c
}

// initializeUser sets up user-related dependencies
func (c *AppContainer) initializeUser() {
	c.UserRepository = repository.NewInMemoryUserRepository()
	c.UserUsecase = usecase.NewUserUsecase(c.UserRepository)
	c.UserHandler = handler.NewUserHandler(c.UserUsecase)
}

// initializeOrg sets up org-related dependencies
func (c *AppContainer) initializeOrg() {
	c.OrgRepository = repository.NewInMemoryOrgRepository()
	c.OrgUsecase = usecase.NewOrgUsecase(c.OrgRepository, c.UserRepository)
	c.OrgHandler = handler.NewOrgHandler(c.OrgUsecase)
}
