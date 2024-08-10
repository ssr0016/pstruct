package postgres

import (
	"bytes"
	"context"
	"fmt"
	"strings"
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

func (r *TaskRepository) Search(ctx context.Context, query *task.SearchTaskQuery) (*task.SearchTaskResult, error) {
	var (
		result = &task.SearchTaskResult{
			Tasks: make([]*task.Task, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
		paramIndex      = 1 // Start with 1 for parameter placeholders
	)

	sql.WriteString(`
		SELECT
			id,
			title,
			description,
			status
		FROM tasks
	`)

	// Handling the title search condition
	if len(query.Title) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("title ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Title+"%")
		paramIndex++
	}

	// Handling the description search condition
	if len(query.Description) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("description ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Description+"%")
		paramIndex++
	}

	// Handling the status search condition
	if len(query.Status) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("status ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Status+"%")
		paramIndex++
	}

	// Add WHERE clause if there are any conditions
	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	// Add ORDER BY clause
	sql.WriteString(" ORDER BY id")

	// Getting the count of total results
	count, err := r.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	// Handling pagination with LIMIT and OFFSET
	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	} else {
		// If per_page is not provided, set a default value (e.g., 20)
		sql.WriteString(" LIMIT 20")
	}

	// Execute the final query
	err = r.DB.SelectContext(ctx, &result.Tasks, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	// Assigning the total count to the result
	result.TotalCount = count

	return result, nil
}

func (r *TaskRepository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := r.DB.GetContext(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
}
