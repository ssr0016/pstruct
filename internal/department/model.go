package department

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidDepartmentName   = errors.New("department.invalid-name", "Invalid name")
	ErrInvalidDeparmentID      = errors.New("department.invalid-id", "Invalid id")
	ErrDepartmentAlreadyExists = errors.New("department.already-exists", "Department already exists")
	ErrDepartmentNotFound      = errors.New("department.not-found", "Department not found")
)

type Department struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type CreateDepartmentCommand struct {
	Name string `json:"name"`
}

type UpdateDepartmentCommand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SearchDepartmentQuery struct {
	Name    string `query:"name"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}

type SearchDepartmentResult struct {
	TotalCount  int           `json:"total_count"`
	Departments []*Department `json:"results"`
	Page        int           `query:"page"`
	PerPage     int           `query:"per_page"`
}

func (d *CreateDepartmentCommand) Validate() error {
	if len(d.Name) == 0 {
		return ErrInvalidDepartmentName
	}

	return nil
}

func (d *UpdateDepartmentCommand) Validate() error {
	if d.ID <= 0 {
		return ErrInvalidDeparmentID
	}

	if len(d.Name) == 0 {
		return ErrInvalidDepartmentName
	}

	return nil
}
