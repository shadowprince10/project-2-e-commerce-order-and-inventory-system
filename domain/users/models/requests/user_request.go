package requests

// RegistrationRequest input model for POST /register
type RegistrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// LoginRequest input model for POST /login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateProfileRequest input model for PUT /profile
type UpdateProfileRequest struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}
