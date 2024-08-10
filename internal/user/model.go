package user

type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

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
