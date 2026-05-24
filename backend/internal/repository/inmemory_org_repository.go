package repository

import (
	"encoding/json"
	"errors"
	"os"
	"salesforce-user-manager/internal/domain"
)

// InMemoryOrgRepository is an in-memory implementation of OrgRepository with JSON persistence
type InMemoryOrgRepository struct {
	orgs     map[string]*domain.Org
	filename string
}

// NewInMemoryOrgRepository creates a new InMemoryOrgRepository with JSON persistence
func NewInMemoryOrgRepository() *InMemoryOrgRepository {
	repo := &InMemoryOrgRepository{
		orgs:     make(map[string]*domain.Org),
		filename: "orgs.json",
	}
	repo.load()
	return repo
}

// load reads orgs from JSON file
func (r *InMemoryOrgRepository) load() {
	file, err := os.Open(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, start with empty
			return
		}
		// For now, ignore other errors, perhaps log later
		return
	}
	defer file.Close()

	var orgs []*domain.Org
	if err := json.NewDecoder(file).Decode(&orgs); err != nil {
		// Ignore decode errors
		return
	}

	for _, org := range orgs {
		r.orgs[org.Id] = org
	}
}

// save writes orgs to JSON file
func (r *InMemoryOrgRepository) save() error {
	file, err := os.Create(r.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var orgs []*domain.Org
	for _, org := range r.orgs {
		orgs = append(orgs, org)
	}

	return json.NewEncoder(file).Encode(orgs)
}

// GetByID retrieves an org by Id
func (r *InMemoryOrgRepository) GetByID(Id string) (*domain.Org, error) {
	org, exists := r.orgs[Id]
	if !exists {
		return nil, errors.New("org not found")
	}
	return org, nil
}

// GetAll retrieves all orgs
func (r *InMemoryOrgRepository) GetAll() ([]*domain.Org, error) {
	var orgs []*domain.Org
	for _, org := range r.orgs {
		orgs = append(orgs, org)
	}
	return orgs, nil
}

// Create creates a new org
func (r *InMemoryOrgRepository) Create(org *domain.Org) error {
	r.orgs[org.Id] = org
	return r.save()
}

// Update updates an existing org
func (r *InMemoryOrgRepository) Update(org *domain.Org) error {
	if _, exists := r.orgs[org.Id]; !exists {
		return errors.New("org not found")
	}
	r.orgs[org.Id] = org
	return r.save()
}

// Delete deletes an org by Id
func (r *InMemoryOrgRepository) Delete(Id string) error {
	if _, exists := r.orgs[Id]; !exists {
		return errors.New("org not found")
	}
	delete(r.orgs, Id)
	return r.save()
}
