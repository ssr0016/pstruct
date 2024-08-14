package user

import (
	"strings"
	"task-management-system/internal/api/errors"
	util "task-management-system/pkg/util/password"
	"task-management-system/pkg/util/validation"
)

var (
	ErrInvalidEmail       = errors.New("user.invalid-email", "Invalid email")
	ErrInvalidID          = errors.New("user.invalid-id", "Invalid id")
	ErrUserAlreadyExists  = errors.New("user.already-exists", "User already exists")
	ErrUserNotFound       = errors.New("user.not-found", "User not found")
	ErrInvalidPassword    = errors.New("user.invalid-password", "Invalid password")
	ErrInvalidFirstName   = errors.New("user.invalid-first-name", "Invalid first name")
	ErrInvalidLastName    = errors.New("user.invalid-last-name", "Invalid last name")
	ErrInvalidAddress     = errors.New("user.invalid-address", "Invalid address")
	ErrInvalidPhoneNumber = errors.New("user.invalid-phone-number", "Invalid phone number")
	ErrInvalidDateOfBirth = errors.New("user.invalid-date-of-birth", "Invalid date of birth")
	ErrEmailAlreadyExists = errors.New("user.email-already-exists", "Email already exists")
)

type User struct {
	ID           int    `db:"id" json:"id"`
	FirstName    string `db:"first_name" json:"first_name"`
	LastName     string `db:"last_name" json:"last_name"`
	Email        string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
	Address      string `db:"address" json:"address"`
	PhoneNumber  string `db:"phone_number" json:"phone_number"`
	DateOfBirth  string `db:"date_of_birth" json:"date_of_birth"`
}

// CreateUserRequest represents the request payload for creating a new user.
type CreateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
}

type UpdateUserRequest struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cmd *CreateUserRequest) Validate() error {
	if len(cmd.FirstName) == 0 {
		return ErrInvalidFirstName
	}

	if len(cmd.FirstName) <= 2 {
		return ErrInvalidFirstName
	}

	if len(cmd.LastName) == 0 {
		return ErrInvalidLastName
	}

	if len(cmd.LastName) <= 2 {
		return ErrInvalidLastName
	}

	if len(cmd.Email) == 0 || !validation.IsValidEmail(cmd.Email) {
		return ErrInvalidEmail
	}

	if len(cmd.Password) == 0 || !util.IsValidPassword(cmd.Password) {
		return ErrInvalidPassword
	}

	if len(cmd.Address) == 0 {
		return ErrInvalidAddress
	}

	if len(cmd.PhoneNumber) == 0 {
		return ErrInvalidPhoneNumber
	}

	if len(cmd.DateOfBirth) == 0 {
		return ErrInvalidDateOfBirth
	}

	return nil
}

func (cmd *UpdateUserRequest) Validate() error {
	if cmd.ID <= 0 {
		return ErrInvalidID
	}

	if len(strings.TrimSpace(cmd.FirstName)) == 0 {
		return ErrInvalidFirstName
	}

	if len(cmd.FirstName) < 2 {
		return ErrInvalidFirstName
	}

	if len(strings.TrimSpace(cmd.LastName)) == 0 {
		return ErrInvalidLastName
	}

	if len(cmd.LastName) < 2 {
		return ErrInvalidLastName
	}

	if len(strings.TrimSpace(cmd.Address)) == 0 {
		return ErrInvalidAddress
	}

	if len(cmd.PhoneNumber) == 0 || !validation.IsValidPhoneNumber(cmd.PhoneNumber) {
		return ErrInvalidPhoneNumber
	}

	if len(cmd.Email) > 0 && !validation.IsValidEmail(cmd.Email) {
		return ErrInvalidEmail
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

	if !validation.IsValidEmail(cmd.Email) {
		return ErrInvalidEmail
	}

	if !util.IsValidPassword(cmd.Password) {
		return ErrInvalidPassword
	}

	return nil
}
