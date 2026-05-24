package tests

import (
	"salesforce-user-manager/internal/domain"
	"salesforce-user-manager/internal/repository"
	"salesforce-user-manager/internal/usecase"
	"testing"
)

func TestUserUsecase_GetUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	usecase := usecase.NewUserUsecase(repo)

	// Create a test user
	user := &domain.User{ID: "1", Username: "testuser", Email: "test@example.com", Name: "Test User", IsActive: true}
	err := repo.Create(user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Test GetUser
	retrieved, err := usecase.GetUser("1")
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	if retrieved.ID != "1" {
		t.Errorf("Expected ID '1', got '%s'", retrieved.ID)
	}
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	usecase := usecase.NewUserUsecase(repo)

	// Create test users
	user1 := &domain.User{ID: "1", Username: "user1", Email: "user1@example.com", Name: "User One", IsActive: true}
	user2 := &domain.User{ID: "2", Username: "user2", Email: "user2@example.com", Name: "User Two", IsActive: false}
	repo.Create(user1)
	repo.Create(user2)

	// Test GetAllUsers
	users, err := usecase.GetAllUsers()
	if err != nil {
		t.Fatalf("Failed to get all users: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestUserUsecase_CreateUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	usecase := usecase.NewUserUsecase(repo)

	user := &domain.User{ID: "1", Username: "newuser", Email: "new@example.com", Name: "New User", IsActive: true}
	err := usecase.CreateUser(user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Verify
	retrieved, err := repo.GetByID("1")
	if err != nil {
		t.Fatalf("Failed to retrieve created user: %v", err)
	}
	if retrieved.Username != "newuser" {
		t.Errorf("Expected username 'newuser', got '%s'", retrieved.Username)
	}
}
