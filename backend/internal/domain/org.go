package domain

// Org represents a Salesforce organization
type Org struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Active         bool   `json:"active"`
	IsSandbox      bool   `json:"is_sandbox"`
	ClientId       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	Domain         string `json:"domain"`
	UsernameSuffix string `json:"username_suffix"`
}
