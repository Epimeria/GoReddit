package model

// APICredentials holds the information from user's .cred file
type APICredentials struct {
	UserAgent       string `json:"user_agent"`
	AccountUsername string `json:"username"`
	AccountPassword string `json:"password"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
}

// APITokenCredentials return struct to give AccessToken, TokenType, etc
type APITokenCredentials struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
