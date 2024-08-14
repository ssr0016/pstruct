package user

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidEmail      = errors.New("user.invalid-email", "Invalid email")
	ErrInvalidID         = errors.New("user.invalid-id", "Invalid id")
	ErrUserAlreadyExists = errors.New("user.already-exists", "User already exists")
	ErrUserNotFound      = errors.New("user.not-found", "User not found")
	ErrInvalidPassword   = errors.New("user.invalid-password", "Invalid password")
	ErrInvalidFirstName  = errors.New("user.invalid-first-name", "Invalid first name")
	ErrInvalidLastName   = errors.New("user.invalid-last-name", "Invalid last name")
)

type User struct {
	ID           int    `db:"id" json:"id"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
}

// CreateUserRequest represents the request payload for creating a new user.
type CreateUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"-"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cmd *CreateUserRequest) Validate() error {
	if len(cmd.FirstName) == 0 {
		return ErrInvalidEmail
	}

	if len(cmd.FirstName) <= 2 {
		return ErrInvalidFirstName
	}

	if len(cmd.LastName) <= 2 {
		return ErrInvalidLastName
	}

	if len(cmd.LastName) == 0 {
		return ErrInvalidLastName
	}

	if len(cmd.Email) == 0 {
		return ErrInvalidEmail
	}

	if len(cmd.Password) == 0 {
		return ErrInvalidPassword
	}

	if len(cmd.Password) <= 6 {
		return ErrInvalidPassword
	}

	return nil
}

func (cmd *LoginUserRequest) Validate() error {
	if len(cmd.Email) == 0 {
		return ErrInvalidEmail
	}

	if len(cmd.Password) == 0 {
		return ErrInvalidPassword
	}

	if len(cmd.Password) <= 6 {
		return ErrInvalidPassword
	}

	return nil
}
