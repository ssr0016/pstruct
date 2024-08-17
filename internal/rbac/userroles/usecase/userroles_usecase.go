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

func (uru *UserRoleUseCase) AssignRoles(ctx context.Context, cmd *userroles.CreateUserRolesCommand) error {
	return uru.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := uru.repo.UserRolesTaken(ctx, 0, cmd.RoleID)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return userroles.ErrUserRolesAlreadyExists
		}

		err = uru.repo.AssignRoles(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (uru *UserRoleUseCase) RemoveUserRoles(ctx context.Context, cmd *userroles.RemoveUserRolesCommand) error {
	return uru.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		if cmd.UserID <= 0 {
			return userroles.ErrInvalidUserID
		}

		if cmd.RoleID <= 0 {
			return userroles.ErrInvalidRoleID
		}

		err := uru.repo.RemoveUserRoles(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (uru *UserRoleUseCase) GetUserRolesByID(ctx context.Context, id int) ([]*userroles.UserRole, error) {
	result, err := uru.repo.GetUserRolesByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, userroles.ErrUserRolesNotFound
	}

	return result, nil
}

func (uru *UserRoleUseCase) UpdateUserRoles(ctx context.Context, cmd *userroles.UpdateUserRolesCommand) error {
	return uru.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		existingRoles, err := uru.repo.GetUserRolesByID(ctx, cmd.UserID)
		if err != nil {
			return err
		}

		if existingRoles == nil {
			return userroles.ErrUserRolesNotFound
		}

		result, err := uru.repo.UserRolesTaken(ctx, cmd.UserID, cmd.RoleID)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			return userroles.ErrUserRolesNotFound
		}

		if len(result) > 1 || (len(result) == 1 && result[0].ID != cmd.RoleID) {
			return userroles.ErrUserRolesAlreadyExists
		}

		err = uru.repo.UpdateUserRoles(ctx, &userroles.UpdateUserRolesCommand{
			ID:     cmd.ID,
			UserID: cmd.UserID,
			RoleID: cmd.RoleID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (uru *UserRoleUseCase) SearchUserRoles(ctx context.Context, query *userroles.SearchUserRolesQuery) (*userroles.SearchUserRolesResult, error) {
	if query.Page <= 0 {
		query.Page = uru.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = uru.cfg.Pagination.PageLimit
	}

	result, err := uru.repo.SearchUserRoles(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}
