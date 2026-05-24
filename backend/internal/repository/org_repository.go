package repository

import "salesforce-user-manager/internal/domain"

// OrgRepository defines the interface for org data operations
type OrgRepository interface {
	GetByID(id string) (*domain.Org, error)
	GetAll() ([]*domain.Org, error)
	Create(org *domain.Org) error
	Update(org *domain.Org) error
	Delete(id string) error
}
