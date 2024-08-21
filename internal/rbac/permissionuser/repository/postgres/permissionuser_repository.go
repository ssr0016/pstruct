package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/permissionuser"

	"go.uber.org/zap"
)

type PermissionUserRepository struct {
	db     db.DB
	logger *zap.Logger
}

func NewPermissionUserRepository(db db.DB) *PermissionUserRepository {
	return &PermissionUserRepository{
		db:     db,
		logger: zap.L().Named("permissionuser.repository"),
	}
}

func (ps *PermissionUserRepository) CreateUserPermission(ctx context.Context, cmd *permissionuser.CreateUserPermissionCommand) error {
	return ps.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
		INSERT INTO user_permissions (
			user_id,
			permission_id,
			action,
			scope
		) VALUES (
		 	$1,
			$2,
			$3,
			$4
		) RETURNING id
		`
		actionStr := strings.Join(cmd.Action, ",")

		var id int
		err := tx.QueryRow(ctx, rawSQL, cmd.UserID, cmd.PermissionID, actionStr, cmd.Scope).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (ps *PermissionUserRepository) DeleteUserPermission(ctx context.Context, id int) error {
	return ps.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
		DELETE FROM user_permissions
		WHERE id = $1
		`
		_, err := tx.Exec(ctx, rawSQL, id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (ps *PermissionUserRepository) GetUserPermissions(ctx context.Context, query *permissionuser.UserPermissionsQuery) (*permissionuser.UserPermissionsResult, error) {
	var (
		result = &permissionuser.UserPermissionsResult{
			UserPermissions: make([]*permissionuser.UserPermission, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
		paramIndex      = 1
	)
	sql.WriteString(`
		SELECT
			id, 
			user_id,
			permission_id,
			action,
			scope
		FROM user_permissions
	`)

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(" ORDER BY id")

	count, err := ps.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = ps.db.Select(ctx, &result.UserPermissions, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil

}

func (ps *PermissionUserRepository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := ps.db.Get(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ps *PermissionUserRepository) GetUserPermissionByID(ctx context.Context, id int) (*permissionuser.UserPermission, error) {
	var result permissionuser.UserPermission

	rawSQL := `
		SELECT 
			id, 
			user_id,
			permission_id,
			action,
			scope
		FROM user_permissions
		WHERE
			id = $1
	`

	err := ps.db.Get(ctx, &result, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &result, nil
}

func (ps *PermissionUserRepository) GetAllUserPermissions(ctx context.Context, userID string) ([]*permissionuser.UserPermission, error) {
	var result []*permissionuser.UserPermission

	// Define the raw SQL query
	rawSQL := `
SELECT id, user_id, permission_id, action, scope
FROM user_permissions
WHERE user_id = (SELECT id FROM users WHERE email = $1)
`

	err := ps.db.Select(ctx, &result, rawSQL, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
