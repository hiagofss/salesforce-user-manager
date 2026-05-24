package repository

import "salesforce-user-manager/internal/domain"

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id string) error
}
