package task

import (
	"task-management-system/internal/api/errors"
)

var (
	ErrInvalidTitle      = errors.New("task.invalid-title", "Invalid title")
	ErrInvalidID         = errors.New("task.invalid-id", "Invalid id")
	ErrTaskAlreadyExists = errors.New("task.already-exists", "Task already exists")
	ErrTaskNotFound      = errors.New("task.not-found", "Task not found")
)

type Task struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Status      string `db:"status" json:"status"`
}

type CreateTaskCommand struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTaskCommand struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SearchTaskQuery struct {
	Title       string `query:"title"`
	Description string `query:"description"`
	Status      string `query:"status"`
	PerPage     int    `query:"per_page"`
	Page        int    `query:"page"`
}

type SearchTaskResult struct {
	TotalCount int     `json:"total_count"`
	Tasks      []*Task `json:"results"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

func (cmd *CreateTaskCommand) Validate() error {
	if len(cmd.Title) == 0 {
		return ErrInvalidTitle
	}
	return nil
}

func (cmd *UpdateTaskCommand) Validate() error {
	if cmd.ID <= 0 {
		return ErrInvalidID
	}

	if len(cmd.Title) == 0 {
		return ErrInvalidTitle
	}

	return nil
}
