package usercase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/permissions"
	"task-management-system/internal/rbac/permissions/repository/postgres"
	"task-management-system/internal/user"

	"go.uber.org/zap"
)

type PermissionUseCase struct {
	repo *postgres.PermissionsRepository
	cfg  *config.Config
	db   db.DB
	log  *zap.Logger
}

func NewPermissionUseCase(db db.DB, cfg *config.Config, user user.Service) permissions.Service {
	return &PermissionUseCase{
		repo: postgres.NewPermissionRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("permission.usecase"),
	}
}

func (pu *PermissionUseCase) CreatePermissions(ctx context.Context, cmd *permissions.CreatePermissionCommand) error {
	return pu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		err := pu.repo.CreatePermissions(ctx, &permissions.CreatePermissionCommand{
			Name: cmd.Name,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (pu *PermissionUseCase) GetPermissions(ctx context.Context) ([]*permissions.Permission, error) {
	result, err := pu.repo.GetUserPermissions(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pu *PermissionUseCase) GetPermissionByID(ctx context.Context, id int) (*permissions.Permission, error) {
	result, err := pu.repo.GetPermissionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, permissions.ErrPermissionNotFound
	}

	return result, nil
}
