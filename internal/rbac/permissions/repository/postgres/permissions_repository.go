package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"
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
		// Convert the slice to a CSV format
		actionsCSV := strings.Join(cmd.Actions, ",")

		rawSQL := `
            INSERT INTO permissions (
                actions
            ) VALUES (
                $1
            ) RETURNING id	
        `

		var id int
		err := tx.QueryRow(ctx, rawSQL, actionsCSV).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (p *PermissionsRepository) GetUserPermissions(ctx context.Context) ([]*permissions.PermissionDTO, error) {
	var result []*permissions.PermissionDTO

	rawSQL := `
		SELECT 
			id,
			actions
		FROM permissions
	`

	err := p.db.Select(ctx, &result, rawSQL)
	if err != nil {
		return nil, err
	}

	// Convert CSV format to slice of strings
	for _, perm := range result {
		// Check if Actions is not empty before splitting
		if perm.Actions != "" {
			perm.Actions = strings.Join(strings.Split(perm.Actions, ","), ",")
		} else {
			perm.Actions = ""
		}
	}

	return result, nil
}
func (p *PermissionsRepository) GetPermissionByID(ctx context.Context, id int) (*permissions.PermissionDTO, error) {
	var result permissions.PermissionDTO

	rawSQL := `
		 SELECT 
		 	id,
			actions
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
