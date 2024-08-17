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
)

type UserRole struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
	RoleID int `db:"role_id" json:"role_id"`
}

type CreateUserRolesCommand struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type UpdateUserRolesCommand struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type RemoveUserRolesCommand struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type SearchUserRolesQuery struct {
	UserID  int `query:"user_id"`
	RoleID  int `query:"role_id"`
	Page    int `query:"page"`
	PerPage int `query:"per_page"`
}

type SearchUserRolesResult struct {
	TotalCount int         `json:"total_count"`
	UserRoles  []*UserRole `json:"results"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
}

func (ur *CreateUserRolesCommand) Validate() error {

	if ur.UserID <= 0 {
		return ErrInvalidUserID
	}

	if ur.RoleID <= 0 {
		return ErrInvalidRoleID
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

func (ur *UpdateUserRolesCommand) Validate() error {
	if ur.ID <= 0 {
		return ErrInvalidUserID
	}

	if ur.UserID <= 0 {
		return ErrInvalidUserID
	}

	if ur.RoleID <= 0 {
		return ErrInvalidRoleID
	}

	return nil
}
