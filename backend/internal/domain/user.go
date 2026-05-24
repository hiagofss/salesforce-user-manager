package domain

// User represents a Salesforce user
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	IsActive      bool   `json:"is_active"`
	LastLoginDate string `json:"last_login_date"`
	OrgId         string `json:"org_id"`
}
