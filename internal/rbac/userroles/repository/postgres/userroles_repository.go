package postgres

import (
	"context"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/userroles"

	"github.com/lib/pq"
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

func (ur *UserRoleRepository) AssignRoles(ctx context.Context, cmd *userroles.CreateUserRolesCommand) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO user_roles (
				user_id,
				role_id,
				role_names
			) VALUES (
				$1,
				$2,
				$3
			) RETURNING id
		`

		var id int
		err := tx.QueryRow(ctx, rawSQL, cmd.UserID, cmd.RoleID, pq.Array(cmd.RoleNames)).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})

}

func (ur *UserRoleRepository) RemoveUserRoles(ctx context.Context, cmd *userroles.RemoveUserRolesCommand) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE FROM
				user_roles
			WHERE
				user_id = $1
			AND
				role_id = $2
		`

		_, err := tx.Exec(ctx, rawSQL, cmd.UserID, cmd.RoleID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (ur *UserRoleRepository) GetUserRolesByID(ctx context.Context, userID string) ([]*userroles.UserRole, error) {
	var result []*userroles.UserRole

	// Define the raw SQL query to fetch roles based on the user's email
	rawSQL := `
SELECT ur.id, ur.user_id, ur.role_id, ur.role_names
FROM user_roles ur
JOIN users u ON ur.user_id = u.id
WHERE u.email = $1
`

	err := ur.db.Select(ctx, &result, rawSQL, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
