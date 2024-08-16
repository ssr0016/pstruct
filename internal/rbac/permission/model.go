package permission

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidPermissionName   = errors.New("permission.invalid-name", "Invalid name")
	ErrInvalidPermissionID     = errors.New("permission.invalid-id", "Invalid id")
	ErrPermissionAlreadyExists = errors.New("permission.already-exists", "Permission already exists")
	ErrPermissionNotFound      = errors.New("permission.not-found", "Permission not found")
)

type Permission struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type CreatePermissionCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdatePermissionCommand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SearchPermissionQuery struct {
	Name        string `query:"name"`
	Description string `query:"description"`
	Page        int    `query:"page"`
	PerPage     int    `query:"per_page"`
}

type SearchPermissionResult struct {
	TotalCount  int           `json:"total_count"`
	Permissions []*Permission `json:"results"`
	Page        int           `json:"page"`
	PerPage     int           `json:"per_page"`
}

func (c *CreatePermissionCommand) Validate() error {
	if len(c.Name) == 0 {
		return ErrInvalidPermissionName
	}

	return nil
}

func (c *UpdatePermissionCommand) Validate() error {
	if c.ID <= 0 {
		return ErrInvalidPermissionID
	}

	if len(c.Name) == 0 {
		return ErrInvalidPermissionName
	}

	return nil
}
