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
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
	RoleID int `db:"role_id" json:"role_id"`
}

type CreateUserRoleCommand struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type CreateRemoveUserRoleCommand struct {
	ID    int         `json:"id"`
	Roles []*UserRole `json:"roles"`
}

type SearchUserRoleQuery struct {
	UserID  int `query:"user_id"`
	RoleID  int `query:"role_id"`
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

type SearchUserRoleResult struct {
	TotalCount int         `json:"total_count"`
	UserRoles  []*UserRole `json:"results"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
}

func (cmd *CreateUserRoleCommand) Validate() error {
	if cmd.UserID <= 0 {
		return ErrInvalidUserID
	}

	if cmd.RoleID <= 0 {
		return ErrInvalidRoleID
	}

	return nil
}

func (cmd *CreateRemoveUserRoleCommand) Validate() error {
	if len(cmd.Roles) == 0 {
		return errors.New("user-role.no-roles-provided", "No roles provided for removal")
	}

	if cmd.ID <= 0 {
		return ErrInvalidUserID
	}

	return nil
}
