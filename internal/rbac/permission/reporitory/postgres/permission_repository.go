package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/permission"

	"go.uber.org/zap"
)

type PermissionRepository struct {
	db     db.DB
	logger *zap.Logger
}

func NewPermissionRepository(db db.DB) *PermissionRepository {
	return &PermissionRepository{
		db:     db,
		logger: zap.L().Named("permission.repository"),
	}
}

func (p *PermissionRepository) CreatePermission(ctx context.Context, cmd *permission.CreatePermissionCommand) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO permissions (
				name,
				description
			) VALUES (
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

func (p *PermissionRepository) GetPermissionByID(ctx context.Context, id int) (*permission.Permission, error) {
	var result permission.Permission

	rawSQL := `
		SELECT
			id,
			name,
			description
		FROM permissions
		WHERE 
			id = $1
	`

	err := p.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &result, nil
}

func (p *PermissionRepository) UpdatePermission(ctx context.Context, cmd *permission.UpdatePermissionCommand) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			UPDATE permissions
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

func (p *PermissionRepository) DeletePermission(ctx context.Context, id int) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE FROM permissions
			WHERE id = $1
		`
		_, err := tx.Exec(ctx, rawSQL, id)
		return err
	})
}

func (p *PermissionRepository) SearchPermission(ctx context.Context, query *permission.SearchPermissionQuery) (*permission.SearchPermissionResult, error) {
	var (
		result = &permission.SearchPermissionResult{
			Permissions: make([]*permission.Permission, 0),
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
		FROM permissions
	`)

	if len(query.Name) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("name ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Name+"%")
		paramIndex++
	}

	if len(query.Description) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("description ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Description+"%")
		paramIndex++
	}

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(" ORDER BY id")

	count, err := p.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = p.db.Select(ctx, &result.Permissions, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (r *PermissionRepository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := r.db.Get(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *PermissionRepository) PermissionTaken(ctx context.Context, id int, name string) ([]*permission.Permission, error) {
	var result []*permission.Permission

	rawSQL := `
		SELECT
			id,
			name,
			description
		FROM permissions
		WHERE
			id = $1 OR 
			name = $2
	`

	err := p.db.Select(ctx, &result, rawSQL, id, name)
	if err != nil {
		return nil, err
	}

	return result, nil
}
