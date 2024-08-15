package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"task-management-system/internal/db"
	rabc "task-management-system/internal/rbac/role"

	"go.uber.org/zap"
)

type RoleRository struct {
	db     db.DB
	logger *zap.Logger
}

func NewRoleRepository(db db.DB) *RoleRository {
	return &RoleRository{
		db:     db,
		logger: zap.L().Named("role.repository"),
	}
}

func (r *RoleRository) CreateRole(ctx context.Context, cmd *rabc.CreateRoleCommand) error {
	return r.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO roles (
				name,
				description
			)VALUES(
				$1,
				$2
			) RETURNING id
		`
		var id int

		err := tx.QueryRow(ctx, rawSQL, cmd.Name, cmd.Description).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleRository) GetRoleByID(ctx context.Context, id int) (*rabc.Role, error) {
	var result rabc.Role

	rawSQL := `
		SELECT
			id,
			name,
			description
		FROM roles
		WHERE 
			id = $1	
	`

	err := r.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func (r *RoleRository) UpdateRole(ctx context.Context, cmd *rabc.UpdateRoleCommand) error {
	return r.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			UPDATE roles
			SET
				name = $1,
				description = $2
			WHERE
				id = $3
		`

		_, err := tx.Exec(ctx, rawSQL, cmd.Name, cmd.Description, cmd.ID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleRository) DeleteRole(ctx context.Context, id int) error {
	return r.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE FROM roles
			WHERE id = $1
		`
		_, err := tx.Exec(ctx, rawSQL, id)
		return err
	})
}

func (r *RoleRository) SearchRole(ctx context.Context, query *rabc.SearchRoleQuery) (*rabc.SearchRoleResult, error) {
	var (
		result = &rabc.SearchRoleResult{
			Roles: make([]*rabc.Role, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
		paramIndex      = 1
	)
	sql.WriteString(`
		SELECT
			id,
			name,
			description
		FROM roles
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

	count, err := r.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = r.db.Select(ctx, &result.Roles, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (r *RoleRository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := r.db.Get(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *RoleRository) RoleTaken(ctx context.Context, id int, name string) ([]*rabc.Role, error) {
	var result []*rabc.Role

	rawSQL := `
		SELECT
			*
		FROM roles
		WHERE
			id = $1 OR
			name = $2
	`

	err := r.db.Select(ctx, &result, rawSQL, id, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}
