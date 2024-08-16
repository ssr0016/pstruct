package postgres

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/userroles"

	"go.uber.org/zap"
)

type UserRoleRepository struct {
	db     db.DB
	logger *zap.Logger
}

func NewUserRoleRepository(db db.DB) *UserRoleRepository {
	return &UserRoleRepository{
		db:     db,
		logger: zap.L().Named("userroles.repository"),
	}
}

func (ur *UserRoleRepository) AssignRoleToUser(ctx context.Context, cmd *userroles.CreateUserRoleCommand) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO user_roles (
				user_id,
				role_id
			) VALUES (
				$1,
				$2
			) Returning id
		`

		var id int
		err := tx.QueryRow(ctx, rawSQL, cmd.UserID, cmd.RoleID).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (ur *UserRoleRepository) RemoveRoleFromUser(ctx context.Context, cmd *userroles.CreateRemoveUserRoleCommand) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE
			FROM 
				user_roles
			WHERE 
				user_id = $1
			AND
				 role_id = $2
	`

		_, err := tx.Exec(ctx, rawSQL, &userroles.CreateRemoveUserRoleCommand{
			ID:    cmd.ID,
			Roles: cmd.Roles,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (ur *UserRoleRepository) SearchUserRole(ctx context.Context, query *userroles.SearchUserRoleQuery) (*userroles.SearchUserRoleResult, error) {
	var (
		result = &userroles.SearchUserRoleResult{
			UserRoles: make([]*userroles.UserRole, 0),
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
			role_id
		FROM user_roles
	`)

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	sql.WriteString(" ORDER BY id")

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	}

	count, err := ur.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (ur *UserRoleRepository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := ur.db.Get(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *UserRoleRepository) UserRoleTaken(ctx context.Context, userID int, roleID int) ([]*userroles.UserRole, error) {
	var result []*userroles.UserRole

	rawSQL := `
		SELECT
			*
		FROM
			user_roles
		WHERE
			user_id = $1 OR 
			role_id = $2	
	`

	err := ur.db.Select(ctx, &result, rawSQL, userID, roleID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
