package permissions

import (
	"fmt"
	"task-management-system/internal/api/errors"
)

var (
	ErrPermissionNotFound = errors.New("permission.not-found", "Permission not found")
	ErrInvalidAction      = errors.New("permission.invalid-action", "Invalid action")
)

type Permission struct {
	ID      int      `db:"id" json:"id"`
	Actions []string `db:"actions" json:"actions"`
}

type PermissionDTO struct {
	ID      int    `db:"id" json:"id"`
	Actions string `db:"actions" json:"actions"`
}

type CreatePermissionCommand struct {
	Actions []string `json:"actions"`
}

func (cmd *CreatePermissionCommand) Validate() error {
	if len(cmd.Actions) == 0 {
		return fmt.Errorf("actions cannot be empty")
	}
	for _, action := range cmd.Actions {
		if !isValidAction(action) {
			return fmt.Errorf("invalid action: %s", action)
		}
	}
	return nil
}

func isValidAction(action string) bool {
	validActions := map[string]bool{
		"read":   true,
		"create": true,
		"delete": true,
		"update": true,
	}
	return validActions[action]
}
