package tests

import (
	"salesforce-user-manager/internal/domain"
	"salesforce-user-manager/internal/repository"
	"salesforce-user-manager/internal/usecase"
	"testing"
)

func TestOrgUsecase_GetOrg(t *testing.T) {
	repo := repository.NewInMemoryOrgRepository()
	usecase := usecase.NewOrgUsecase(repo)

	// Create a test org
	org := &domain.Org{
		ID:             "1",
		Name:           "Test Org",
		Active:         true,
		ClientId:       "client123",
		ClientSecret:   "secret123",
		Domain:         "test.salesforce.com",
		UsernameSuffix: "@test.com",
	}
	err := repo.Create(org)
	if err != nil {
		t.Fatalf("Failed to create org: %v", err)
	}

	// Test GetOrg
	retrieved, err := usecase.GetOrg("1")
	if err != nil {
		t.Fatalf("Failed to get org: %v", err)
	}
	if retrieved.ID != "1" {
		t.Errorf("Expected ID '1', got '%s'", retrieved.ID)
	}
	if retrieved.Name != "Test Org" {
		t.Errorf("Expected Name 'Test Org', got '%s'", retrieved.Name)
	}
}

func TestOrgUsecase_GetAllOrgs(t *testing.T) {
	repo := repository.NewInMemoryOrgRepository()
	usecase := usecase.NewOrgUsecase(repo)

	// Create test orgs
	org1 := &domain.Org{
		ID:             "1",
		Name:           "Org One",
		Active:         true,
		ClientId:       "client1",
		ClientSecret:   "secret1",
		Domain:         "org1.salesforce.com",
		UsernameSuffix: "@org1.com",
	}
	org2 := &domain.Org{
		ID:             "2",
		Name:           "Org Two",
		Active:         false,
		ClientId:       "client2",
		ClientSecret:   "secret2",
		Domain:         "org2.salesforce.com",
		UsernameSuffix: "@org2.com",
	}
	repo.Create(org1)
	repo.Create(org2)

	// Test GetAllOrgs
	orgs, err := usecase.GetAllOrgs()
	if err != nil {
		t.Fatalf("Failed to get all orgs: %v", err)
	}
	if len(orgs) != 2 {
		t.Errorf("Expected 2 orgs, got %d", len(orgs))
	}
}

func TestOrgUsecase_CreateOrg(t *testing.T) {
	repo := repository.NewInMemoryOrgRepository()
	usecase := usecase.NewOrgUsecase(repo)

	org := &domain.Org{
		ID:             "1",
		Name:           "New Org",
		Active:         true,
		ClientId:       "newclient",
		ClientSecret:   "newsecret",
		Domain:         "neworg.salesforce.com",
		UsernameSuffix: "@neworg.com",
	}
	err := usecase.CreateOrg(org)
	if err != nil {
		t.Fatalf("Failed to create org: %v", err)
	}

	// Verify
	retrieved, err := repo.GetByID("1")
	if err != nil {
		t.Fatalf("Failed to retrieve created org: %v", err)
	}
	if retrieved.Name != "New Org" {
		t.Errorf("Expected name 'New Org', got '%s'", retrieved.Name)
	}
}