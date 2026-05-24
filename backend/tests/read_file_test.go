package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"salesforce-user-manager/internal/domain"
	"testing"
)

func TestRunFuncOrg(t *testing.T) {
	fmt.Println("RUNNING TEST")

	// Decode data into array of domain.User
	var orgs []domain.Org
	err := decodeFromFile("../orgs.json", &orgs)
	if err != nil {
		t.Fatalf("Failed to decode orgs: %v", err)
	}

	fmt.Printf("Decoded %d orgs:\n", len(orgs))
	for i, user := range orgs {
		fmt.Printf("User %d: %+v\n", i+1, user)
	}

	// Verify the data
	if len(orgs) != 2 {
		t.Errorf("Expected 2 orgs, got %d", len(orgs))
	}

	// Check first user
	if orgs[0].ID != "12345" || orgs[0].Name != "Sandbox 1" {
		t.Errorf("First user data incorrect: %+v", orgs[0])
	}
}
func TestRunFuncUser(t *testing.T) {
	fmt.Println("RUNNING TEST")

	// Decode data into array of domain.User
	var users []domain.User
	err := decodeFromFile("../users.json", &users)
	if err != nil {
		t.Fatalf("Failed to decode users: %v", err)
	}

	fmt.Printf("Decoded %d users:\n", len(users))
	for i, user := range users {
		fmt.Printf("User %d: %+v\n", i+1, user)
	}

	// Verify the data
	if len(users) != 3 {
		t.Errorf("Expected 3 users, got %d", len(users))
	}

	// Check first user
	if users[0].ID != "1" || users[0].Username != "john.doe" {
		t.Errorf("First user data incorrect: %+v", users[0])
	}
}

func decodeFromFile(fileName string, object any) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(file, object)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}
