package permissionuser

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidUserID          = errors.New("user-role.invalid-user-id", "Invalid user id")
	ErrInvalidPermissionID    = errors.New("user-role.invalid-permission-id", "Invalid permission id")
	ErrInvalidAction          = errors.New("user-role.invalid-action", "Invalid action")
	ErrUserPermissionNotFound = errors.New("user-role.user-permission-not-found", "User permission not found")
	ErrActionIsEmpty          = errors.New("user-role.action-is-empty", "Action is empty")
	ErrScopeIsEmpty           = errors.New("user-role.scope-is-empty", "Scope is empty")
)

type UserPermission struct {
	ID           int    `db:"id" json:"id"`
	UserID       int    `db:"user_id" json:"user_id"`
	PermissionID int    `db:"permission_id" json:"permission_id"`
	Action       string `db:"action" json:"action"`
	Scope        string `db:"scope" json:"scope"`
}

type CreateUserPermissionCommand struct {
	UserID       int      `json:"user_id"`
	PermissionID int      `json:"permission_id"`
	Action       []string `json:"action"`
	Scope        string   `json:"scope"`
}

type UserPermissionsQuery struct {
	UserID       int      `query:"user_id"`
	PermissionID int      `query:"permission_id"`
	Action       []string `query:"action"`
	Scope        string   `query:"scope"`
	Page         int      `query:"page"`
	PerPage      int      `query:"per_page"`
}

type UserPermissionsResult struct {
	TotalCount      int               `json:"total_count"`
	UserPermissions []*UserPermission `json:"results"`
	Page            int               `json:"page"`
	PerPage         int               `json:"per_page"`
}

func (cmd *CreateUserPermissionCommand) Validate() error {
	if cmd.UserID == 0 {
		return ErrInvalidUserID
	}
	if cmd.PermissionID == 0 {
		return ErrInvalidPermissionID
	}

	if len(cmd.Action) == 0 {
		return ErrActionIsEmpty

	}

	return nil
}
