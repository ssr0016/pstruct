package permission

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidUserID = errors.New("permission.invalid-user-id", "Invalid user id")
	ErrInvalidAction = errors.New("permission.invalid-action", "Invalid action")
)

type UserPermission struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"user_id"`
	Action string `db:"action" json:"action"`
	Scope  string `db:"scope" json:"scope"`
}

type AddPermissionCommand struct {
	UserID     int                       `json:"user_id"`
	Permission []*ActionScopesPermission `json:"permission"`
}

type ActionScopesPermission struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`
}

func (cmd *AddPermissionCommand) Validate() error {
	if cmd.UserID <= 0 {
		return ErrInvalidUserID
	}

	if len(cmd.Permission) == 0 {
		return ErrInvalidAction
	}

	return nil
}
