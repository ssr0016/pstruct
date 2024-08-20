package permissions

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidName        = errors.New("permission.invalid-name", "Invalid name")
	ErrPermissionNotFound = errors.New("permission.not-found", "Permission not found")
)

type Permission struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type CreatePermissionCommand struct {
	Name string `json:"name"`
}

func (cmd *CreatePermissionCommand) Validate() error {
	if len(cmd.Name) == 0 {
		return ErrInvalidName
	}

	return nil
}
