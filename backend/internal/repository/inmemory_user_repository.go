package repository

import (
	"encoding/json"
	"errors"
	"os"
	"salesforce-user-manager/internal/domain"
)

// InMemoryUserRepository is an in-memory implementation of UserRepository with JSON persistence
type InMemoryUserRepository struct {
	users    map[string]*domain.User
	filename string
}

// NewInMemoryUserRepository creates a new InMemoryUserRepository with JSON persistence
func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{
		users:    make(map[string]*domain.User),
		filename: "users.json",
	}
	repo.load()
	return repo
}

// load reads users from JSON file
func (r *InMemoryUserRepository) load() {
	file, err := os.Open(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, start with empty
			return
		}
		// For now, ignore other errors
		return
	}
	defer file.Close()

	var users []*domain.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		// Ignore decode errors
		return
	}

	for _, user := range users {
		r.users[user.ID] = user
	}
}

// save writes users to JSON file
func (r *InMemoryUserRepository) save() error {
	file, err := os.Create(r.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var users []*domain.User
	for _, user := range r.users {
		users = append(users, user)
	}

	return json.NewEncoder(file).Encode(users)
}

// GetByID retrieves a user by ID
func (r *InMemoryUserRepository) GetByID(id string) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetAll retrieves all users
func (r *InMemoryUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// Create creates a new user
func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.users[user.ID] = user
	return r.save()
}

// Update updates an existing user
func (r *InMemoryUserRepository) Update(user *domain.User) error {
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return r.save()
}

// Delete deletes a user by ID
func (r *InMemoryUserRepository) Delete(id string) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return r.save()
}

func readJSON(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	datas := []map[string]interface{}{}

	file, _ := os.ReadFile(fileName)
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		// Do some filtering
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}
