package responses

// ProfileResponse output model for successful POST /register and GET /profile
type ProfileResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// LoginResponse output model for POST /login
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse for standardized error output
type ErrorResponse struct {
	Message string `json:"message"`
}
