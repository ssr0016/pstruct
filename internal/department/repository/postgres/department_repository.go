package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"task-management-system/internal/db"
	"task-management-system/internal/department"

	"go.uber.org/zap"
)

type DepartmentRepository struct {
	db     db.DB
	logger *zap.Logger
}

func NewDepartmentRepository(db db.DB) *DepartmentRepository {
	return &DepartmentRepository{
		db:     db,
		logger: zap.L().Named("department.repository"),
	}
}

func (d *DepartmentRepository) CreateDepartment(ctx context.Context, cmd *department.CreateDepartmentCommand) error {
	return d.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO departments (
				name
			) VALUES (
				$1
			) RETURNING id
		`
		var id int

		err := tx.QueryRow(ctx, rawSQL, cmd.Name).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (d *DepartmentRepository) GetDepartmentByID(ctx context.Context, id int) (*department.Department, error) {
	var result department.Department

	rawSQL := `
		SELECT
			id,
			name
		FROM
			departments
		WHERE
			id = $1
	`

	err := d.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func (d *DepartmentRepository) UpdateDepartment(ctx context.Context, cmd *department.UpdateDepartmentCommand) error {
	return d.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			UPDATE departments
			SET
				name = $1	
			WHERE
				id = $2
		`

		_, err := tx.Exec(ctx, rawSQL, cmd.Name, cmd.ID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (d *DepartmentRepository) DeleteDepartment(ctx context.Context, id int) error {
	return d.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE FROM departments
			WHERE id = $1
		`
		_, err := tx.Exec(ctx, rawSQL, id)
		return err
	})
}

func (d *DepartmentRepository) SearchDepartment(ctx context.Context, query *department.SearchDepartmentQuery) (*department.SearchDepartmentResult, error) {
	var (
		result = &department.SearchDepartmentResult{
			Departments: make([]*department.Department, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
		paramIndex      = 1
	)

	sql.WriteString(`
		SELECT
			id,
			name
		FROM departments
	`)

	if len(query.Name) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("name ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Name+"%")
		paramIndex++
	}

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(" ORDER BY id")

	// Getting the count of total results
	count, err := d.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = d.db.Select(ctx, &result.Departments, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (d *DepartmentRepository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := d.db.Get(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *DepartmentRepository) DepartmentTaken(ctx context.Context, id int, name string) ([]*department.Department, error) {
	var result []*department.Department

	rawSQL := `
		SELECT
			*
		FROM departments
		WHERE
			id = $1 OR
			name = $2
	`

	err := d.db.Select(ctx, &result, rawSQL, id, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}
