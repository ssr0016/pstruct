package postgres

import (
	"context"
	"task-management-system/internal/db"
	permission "task-management-system/internal/rbac/userpermission"

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

func (p *PermissionRepository) Add(ctx context.Context, entities []*permission.UserPermission) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		// Prepare the SQL statement
		rawSQL := `
			INSERT INTO user_permissions (
				user_id,
				action,
				scope
			) VALUES (
				$1,
				$2,
				$3
			) RETURNING id
		`

		var id int

		for _, entity := range entities {
			err := tx.QueryRow(ctx, rawSQL, entity.UserID, entity.Action, entity.Scope).Scan(&id)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *PermissionRepository) Delete(ctx context.Context, userID int) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE FROM user_permissions
			WHERE 
				user_id = $1
		`
		_, err := tx.Exec(ctx, rawSQL, userID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (p *PermissionRepository) GetActionByUserID(ctx context.Context, userID int) ([]*string, error) {
	var result []*string

	rawSQL := `
		SELECT
			action
		FROM user_permissions
		WHERE
			user_id = $1
	`

	err := p.db.Select(ctx, &result, rawSQL, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PermissionRepository) GetListPermission(ctx context.Context, userID int) ([]*permission.ActionScopesPermission, error) {
	var result []*permission.ActionScopesPermission

	rawSQL := `
		SELECT
			action,
			scope
		FROM user_permissions
		WHERE
			user_id = $1
	`

	err := p.db.Select(ctx, &result, rawSQL, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
