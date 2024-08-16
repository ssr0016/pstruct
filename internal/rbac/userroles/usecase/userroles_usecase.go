package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/userroles"
	"task-management-system/internal/rbac/userroles/repository/postgres"

	"go.uber.org/zap"
)

type UserRoleUseCase struct {
	repo *postgres.UserRoleRepository
	db   db.DB
	cfg  *config.Config
	log  *zap.Logger
}

func NewUserRoleUseCase(db db.DB, cfg *config.Config) userroles.Service {
	return &UserRoleUseCase{
		repo: postgres.NewUserRoleRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("userroles.usecase"),
	}
}

// AssignRoleToUser implements userroles.Service.
func (ur *UserRoleUseCase) AssignRoleToUser(ctx context.Context, cmd *userroles.CreateUserRoleCommand) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := ur.repo.UserRoleTaken(ctx, cmd.UserID, cmd.RoleID)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return userroles.ErrUserRoleAlreadyExists
		}

		err = ur.repo.AssignRoleToUser(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

// RemoveRoleFromUser implements userroles.Service.
func (ur *UserRoleUseCase) RemoveRoleFromUser(ctx context.Context, cmd *userroles.CreateRemoveUserRoleCommand) error {
	return ur.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		err := ur.repo.RemoveRoleFromUser(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

// SearchUserRole implements userroles.Service.
func (ur *UserRoleUseCase) SearchUserRole(ctx context.Context, query *userroles.SearchUserRoleQuery) (*userroles.SearchUserRoleResult, error) {
	if query.Page <= 0 {
		query.Page = ur.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = ur.cfg.Pagination.PageLimit
	}

	result, err := ur.repo.SearchUserRole(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}
