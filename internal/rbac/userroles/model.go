package userroles

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidUserID         = errors.New("user-role.invalid-user-id", "Invalid user id")
	ErrInvalidRoleID         = errors.New("user-role.invalid-role-id", "Invalid role id")
	ErrUserRoleNotFound      = errors.New("user-role.not-found", "User role not found")
	ErrUserRoleAlreadyExists = errors.New("user-role.already-exists", "User role already exists")
)

type UserRole struct {
	UserID int `db:"user_id" json:"user_id"`
	RoleID int `db:"role_id" json:"role_id"`
}

func (ur *UserRole) Validate() error {

	if ur.UserID <= 0 {
		return ErrInvalidUserID
	}

	if ur.RoleID <= 0 {
		return ErrInvalidRoleID
	}

	return nil
}
