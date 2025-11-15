package responses

import (
	"encoding/json"
	"net/http"

	"github.com/shadowprince10/project-2-e-commerce-order-and-inventory-system/domain/users/entities"
)

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

// mapUserToResponse function converts a domain User entity to public ProfileResponse Data Transfer Object for better security, decoupling, readability, and validation.
func mapUserToResponse(u entities.User) ProfileResponse {
	return ProfileResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
}

// respondToError function writes a standardized JSON error response
func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: msg})
}

// respondWithJSON function writes a standardized JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
