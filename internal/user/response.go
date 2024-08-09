package user

// UserResponse represents the response payload for user details.
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// ErrorResponse represents the response payload for error messages.
type ErrorResponse struct {
	Error string `json:"error"`
}
