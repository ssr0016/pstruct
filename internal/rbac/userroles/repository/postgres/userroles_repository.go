package postgres

import (
	"context"
	"task-management-system/internal/db"

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

func (ur *UserRoleRepository) Assign(ctx context.Context, userID, roleID int) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO user_roles (user_id, role_id)
			VALUES ($1, $2)
		`
		_, err := tx.Exec(ctx, rawSQL, userID, roleID)

		return err
	})
}
