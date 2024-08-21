package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	role "task-management-system/internal/rbac/role/repository/postgres"
	"task-management-system/internal/rbac/userroles"
	"task-management-system/internal/rbac/userroles/repository/postgres"

	"go.uber.org/zap"
)

type UserRoleUseCase struct {
	repo     *postgres.UserRoleRepository
	roleRepo *role.RoleRository
	db       db.DB
	cfg      *config.Config
	log      *zap.Logger
}

func NewUserRoleUseCase(db db.DB, cfg *config.Config) userroles.Service {
	return &UserRoleUseCase{
		repo:     postgres.NewUserRoleRepository(db),
		roleRepo: role.NewRoleRepository(db),
		db:       db,
		cfg:      cfg,
		log:      zap.L().Named("userroles.usecase"),
	}
}

func (uru *UserRoleUseCase) AssignRoles(ctx context.Context, cmd *userroles.CreateUserRolesCommand) error {
	return uru.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		err := uru.repo.AssignRoles(ctx, cmd)
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

func (uru *UserRoleUseCase) GetUserRolesByID(ctx context.Context, userID string) ([]*userroles.UserRole, error) {
	result, err := uru.repo.GetUserRolesByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, userroles.ErrUserNotFound
	}

	return result, nil
}

func (uru *UserRoleUseCase) GetRoleNameByID(ctx context.Context, roleID int) (string, error) {
	role, err := uru.roleRepo.GetRoleByID(ctx, roleID)
	if err != nil {
		return "", err
	}

	return role.Name, nil
}
