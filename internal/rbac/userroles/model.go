package userroles

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidUserID          = errors.New("user-role.invalid-user-id", "Invalid user id")
	ErrInvalidRoleID          = errors.New("user-role.invalid-role-id", "Invalid role id")
	ErrUserRolesAlreadyExists = errors.New("user-role.already-exists", "User role already exists")
	ErrUserIDNotFound         = errors.New("user-role.user-not-found", "User not found")
	ErrUserRolesNotFound      = errors.New("user-roles.not-found", "User roles not found")
	ErrUserNotFound           = errors.New("user-role.user-not-found", "User not found")
)

type UserRole struct {
	ID        int    `db:"id" json:"id"`
	UserID    int    `db:"user_id" json:"user_id"`
	RoleID    int    `db:"role_id" json:"role_id"`
	RoleNames string `db:"role_names" json:"role_names"`
}

type CreateUserRolesCommand struct {
	UserID    string   `json:"user_id"`
	RoleID    int      `json:"role_id"`
	RoleNames []string `json:"role_names"`
}

type RemoveUserRolesCommand struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

func (ur *CreateUserRolesCommand) Validate() error {

	if ur.UserID == "" {
		return ErrInvalidUserID
	}

	if ur.RoleID <= 0 {
		return ErrInvalidRoleID
	}

	if len(ur.RoleNames) == 0 {
		return errors.New("user-role.no-role-names", "Role names are required")
	}

	return nil
}

func (ur *RemoveUserRolesCommand) Validate() error {
	if ur.UserID <= 0 {
		return ErrInvalidUserID
	}

	if ur.RoleID <= 0 {
		return ErrInvalidRoleID
	}

	return nil
}
