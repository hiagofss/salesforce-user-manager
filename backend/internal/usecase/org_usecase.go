package usecase

import (
	"errors"
	"fmt"
	"salesforce-user-manager/internal/domain"
	"salesforce-user-manager/internal/repository"
	"salesforce-user-manager/internal/service"
)

// OrgUsecase handles business logic for orgs
type OrgUsecase struct {
	orgRepo  repository.OrgRepository
	userRepo repository.UserRepository
}

// NewOrgUsecase creates a new OrgUsecase
func NewOrgUsecase(orgRepo repository.OrgRepository, userRepo repository.UserRepository) *OrgUsecase {
	return &OrgUsecase{orgRepo: orgRepo, userRepo: userRepo}
}

// GetOrg retrieves an org by ID
func (u *OrgUsecase) GetOrg(id string) (*domain.Org, error) {
	return u.orgRepo.GetByID(id)
}

// GetAllOrgs retrieves all orgs
func (u *OrgUsecase) GetAllOrgs() ([]*domain.Org, error) {
	return u.orgRepo.GetAll()
}

// CreateOrg creates a new org
func (u *OrgUsecase) CreateOrg(org *domain.Org) (*domain.Org, error) {

	if org.ClientId == "" || org.ClientSecret == "" || org.Domain == "" || org.UsernameSuffix == "" {
		return nil, errors.New("Missing required fields")
	}

	sfService := service.NewSalesforceService(org.Domain, org.ClientId, org.ClientSecret)

	response, err := sfService.Get("/services/data/v66.0/query?q=SELECT+Id,Name,OrganizationType,InstanceName,IsSandbox+FROM+Organization")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if orgRecords, ok := response["records"].([]interface{}); ok {
		if len(orgRecords) == 1 {
			if orgRecord, ok := orgRecords[0].(map[string]interface{}); ok {
				org.IsSandbox = getBool(orgRecord, "IsSandbox")
				org.Active = getBool(orgRecord, "IsActive")
				org.Name = getString(orgRecord, "Name")
				org.Id = getString(orgRecord, "Id")
			}
		} else {
			return nil, errors.New("More than one Org found")
		}
	} else {
		return nil, errors.New("Failed to get org")
	}

	orgExistent, err := u.orgRepo.GetByID(org.Id)
	if err != nil && err.Error() != "org not found" {
		return nil, err
	}

	if orgExistent != nil {
		u.orgRepo.Update(org)
	}

	if orgExistent == nil {
		u.orgRepo.Create(org)
	}

	return org, nil
}

// UpdateOrg updates an existing org
func (u *OrgUsecase) UpdateOrg(org *domain.Org) error {
	return u.orgRepo.Update(org)
}

// DeleteOrg deletes an org by ID
func (u *OrgUsecase) DeleteOrg(id string) error {
	return u.orgRepo.Delete(id)
}

// SyncUsersFromSalesforce fetches users from Salesforce for the given org
func (u *OrgUsecase) SyncUsersFromSalesforce(orgID string) error {
	org, err := u.orgRepo.GetByID(orgID)
	if err != nil {
		return err
	}

	sfService := service.NewSalesforceService(org.Domain, org.ClientId, org.ClientSecret)

	var orgQueryResponse []map[string]interface{}

	resp, err := sfService.Get("/services/data/v66.0/query?q=SELECT+Id,Name,OrganizationType,InstanceName,IsSandbox+FROM+Organization")
	if err != nil {
		return err
	}

	fmt.Println(resp)

	if orgRecords, ok := resp["records"].([]interface{}); ok {
		for _, r := range orgRecords {
			if rec, ok := r.(map[string]interface{}); ok {
				orgQueryResponse = append(orgQueryResponse, rec)
			}
		}
	}

	// Note: In production, handle authentication errors properly
	// For now, assuming authenticate is called internally if needed

	var allRecords []map[string]interface{}
	nextUrl := "/services/data/v66.0/query?q=SELECT+Id,Username,Email,Name,IsActive,LastLoginDate+FROM+User+WHERE+UserType='Standard'"

	for {
		resp, err := sfService.Get(nextUrl)
		if err != nil {
			return err
		}

		if records, ok := resp["records"].([]interface{}); ok {
			for _, r := range records {
				if rec, ok := r.(map[string]interface{}); ok {
					allRecords = append(allRecords, rec)
				}
			}
		}

		if next, ok := resp["nextRecordsUrl"].(string); ok && next != "" {
			nextUrl = next
		} else {
			break
		}
	}

	// Save users to repository
	for _, rec := range allRecords {
		user := &domain.User{
			ID:            getString(rec, "Id"),
			Username:      getString(rec, "Username"),
			Email:         getString(rec, "Email"),
			Name:          getString(rec, "Name"),
			IsActive:      getBool(rec, "IsActive"),
			LastLoginDate: getString(rec, "LastLoginDate"),
			OrgId:         orgID,
		}

		// Save to user repository
		if err := u.userRepo.Create(user); err != nil {
			// Log error but continue
			continue
		}
	}

	return nil
}

// Helper functions to safely extract values
func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok && val != nil {
		if s, ok := val.(string); ok {
			return s
		}
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if val, ok := m[key]; ok && val != nil {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}
