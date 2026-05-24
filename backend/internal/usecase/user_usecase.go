package usecase

import (
	"salesforce-user-manager/internal/domain"
	"salesforce-user-manager/internal/repository"
)

// UserUsecase handles business logic for users
type UserUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase creates a new UserUsecase
func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// GetUser retrieves a user by ID
func (u *UserUsecase) GetUser(id string) (*domain.User, error) {
	return u.repo.GetByID(id)
}

// GetAllUsers retrieves all users
func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	return u.repo.GetAll()
}

// CreateUser creates a new user
func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.repo.Create(user)
}

// UpdateUser updates an existing user
func (u *UserUsecase) UpdateUser(user *domain.User) error {
	return u.repo.Update(user)
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(id string) error {
	return u.repo.Delete(id)
}
