package rabc

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrNameIsEmpty       = errors.New("role.name-is-empty", "Name is empty")
	ErrNamesIsEmpty      = errors.New("role.names-is-empty", "Names is empty")
	ErrInvalidID         = errors.New("role.invalid-id", "Invalid id")
	ErrRoleAlreadyExists = errors.New("role.already-exists", "Role already exists")
	ErrRoleNotFound      = errors.New("role.not-found", "Role not found")
)

type Role struct {
	ID          int      `db:"id" json:"id"`
	Name        []string `db:"name" json:"name"`
	Description string   `db:"description" json:"description"`
}

type RoleDTO struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type CreateRoleCommand struct {
	Name        []string `json:"name"`
	Description string   `json:"description"`
}

type UpdateRoleCommand struct {
	ID          int      `json:"id"`
	Name        []string `json:"name"`
	Description string   `json:"description"`
}

type SearchRoleQuery struct {
	Name        string `query:"name"`
	Description string `query:"description"`
	Page        int    `query:"page"`
	PerPage     int    `query:"per_page"`
}

type SearchRoleResult struct {
	TotalCount int     `json:"total_count"`
	Roles      []*Role `json:"results"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

func (r *CreateRoleCommand) Validate() error {
	if len(r.Name) == 0 {
		return ErrNameIsEmpty
	}

	for _, name := range r.Name {
		if len(name) == 0 {
			return ErrNameIsEmpty
		}

		if inValidNames(name) {
			return ErrNamesIsEmpty
		}
	}

	return nil
}

func inValidNames(name string) bool {
	validNames := map[string]bool{
		"admin":   true,
		"manager": true,
		"hr":      true,
	}

	return !validNames[name]
}

func (r *UpdateRoleCommand) Validate() error {
	if r.ID <= 0 {
		return ErrInvalidID
	}

	if len(r.Name) == 0 {
		return ErrNameIsEmpty
	}

	return nil
}
