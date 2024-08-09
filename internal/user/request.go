package user

// CreateUserRequest represents the request payload for creating a new user.
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateUserRequest represents the request payload for updating an existing user.
type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
