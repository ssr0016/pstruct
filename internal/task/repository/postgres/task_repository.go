package postgres

import (
	"context"
	"task-management-system/internal/task"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, cmd *task.CreateTaskCommand) error {
	rawSQL := `
		INSERT INTO tasks ( 
		title,
		description,
		status
		) VALUES (
		 $1,
		 $2,
		$3
		) RETURNING id
	`
	var id int

	err := r.DB.QueryRowxContext(ctx, rawSQL, cmd.Title, cmd.Description, cmd.Status).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id int) (*task.Task, error) {
	var result task.Task

	rawSQL := `
		SELECT
			id,
			title,
			description,
			status
		FROM tasks
		WHERE 
			id = $1	
	`

	err := r.DB.Get(&result, rawSQL, id)
	return &result, err
}

func (r *TaskRepository) Update(ctx context.Context, cmd *task.UpdateTaskCommand) error {
	rawSQL := `
		UPDATE tasks
		SET
			title = $1,
			description = $2,
			status = $3
		WHERE
			id = $4
	`
	_, err := r.DB.ExecContext(ctx, rawSQL, cmd.Title, cmd.Description, cmd.Status, cmd.ID)
	return err
}

func (r *TaskRepository) Delete(ctx context.Context, id int) error {
	rawSQL := `
		DELETE FROM tasks
		WHERE id = $1
	`
	_, err := r.DB.ExecContext(ctx, rawSQL, id)
	return err
}

func (r *TaskRepository) GetAll() ([]task.Task, error) {
	var tasks []task.Task
	err := r.DB.Select(&tasks, "SELECT * FROM tasks")
	return tasks, err
}
