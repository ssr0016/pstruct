package postgres

import (
	"context"
	"database/sql"
	"errors"
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
		namesCSV := strings.Join(cmd.Name, ", ")

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

		err := tx.QueryRow(ctx, rawSQL, namesCSV, cmd.Description).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleRository) GetRoleByID(ctx context.Context, id int) (*rabc.RoleDTO, error) {
	var result rabc.RoleDTO

	rawSQL := `
		SELECT
			id,
			name,
			description
		FROM roles
		WHERE id = $1
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

func (r *RoleRository) GetRoles(ctx context.Context) ([]*rabc.RoleDTO, error) {
	var result []*rabc.RoleDTO

	rawSQL := `
		SELECT
			id,
			name,
			description
		FROM roles
	`

	err := r.db.Select(ctx, &result, rawSQL)
	if err != nil {
		return nil, err
	}

	for _, role := range result {
		// Convert CSV format to slice of strings
		if role.Name != "" {
			role.Name = strings.Join(strings.Split(role.Name, ","), ",")
		} else {
			role.Name = ""
		}
	}

	return result, nil
}

func (r *RoleRository) getRoleNameByID(ctx context.Context, roleID int) (string, error) {
	var roleName string

	rawSQL := `
		SELECT
			name
		FROM roles
		WHERE 
			id = $1
	`

	err := r.db.Get(ctx, &roleName, rawSQL, roleID)
	if err != nil {
		return "", err
	}

	return roleName, nil
}
