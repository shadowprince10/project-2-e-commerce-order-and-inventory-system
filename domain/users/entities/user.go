package entities

// User struct represents the core user entity
type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PasswordHashed  string `json:"-"` // hidden from JSON output
	Phone           string `json:"phone"`
	IsAuthenticated bool   `json:"-"` // to track user status (authenticated or not)
}
