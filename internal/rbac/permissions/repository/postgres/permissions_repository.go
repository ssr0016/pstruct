package postgres

import (
	"context"
	"database/sql"
	"errors"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/permissions"

	"go.uber.org/zap"
)

type PermissionsRepository struct {
	db     db.DB
	logger *zap.Logger
}

func NewPermissionRepository(db db.DB) *PermissionsRepository {
	return &PermissionsRepository{
		db:     db,
		logger: zap.L().Named("permissions.repository"),
	}
}

func (p *PermissionsRepository) CreatePermissions(ctx context.Context, cmd *permissions.CreatePermissionCommand) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO permissions (
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

func (p *PermissionsRepository) GetUserPermissions(ctx context.Context) ([]*permissions.Permission, error) {
	var result []*permissions.Permission

	rawSQL := `
		 SELECT 
		 	id,
			name
		FROM permissions
	`

	err := p.db.Select(ctx, &result, rawSQL)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PermissionsRepository) GetPermissionByID(ctx context.Context, id int) (*permissions.Permission, error) {
	var result permissions.Permission

	rawSQL := `
		 SELECT 
		 	id,
			name
		FROM permissions
		WHERE id = $1
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
