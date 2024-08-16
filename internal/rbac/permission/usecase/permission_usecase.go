package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/permission"
	"task-management-system/internal/rbac/permission/reporitory/postgres"

	"go.uber.org/zap"
)

type PermissionUseCase struct {
	repo *postgres.PermissionRepository
	cfg  *config.Config
	db   db.DB
	log  *zap.Logger
}

func NewPermissionUseCase(db db.DB, cfg *config.Config) permission.Service {
	return &PermissionUseCase{
		repo: postgres.NewPermissionRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("permission.usecase"),
	}
}

// CreatePermission implements permission.Service.
func (pu *PermissionUseCase) CreatePermission(ctx context.Context, cmd *permission.CreatePermissionCommand) error {
	return pu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := pu.repo.PermissionTaken(ctx, 0, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return permission.ErrPermissionAlreadyExists
		}

		err = pu.repo.CreatePermission(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

// DeletePermission implements permission.Service.
func (pu *PermissionUseCase) DeletePermission(ctx context.Context, id int) error {
	return pu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := pu.repo.GetPermissionByID(ctx, id)
		if err != nil {
			return err
		}

		if result == nil {
			return permission.ErrPermissionNotFound
		}

		err = pu.repo.DeletePermission(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

// GetPermissionByID implements permission.Service.
func (pu *PermissionUseCase) GetPermissionByID(ctx context.Context, id int) (*permission.Permission, error) {
	result, err := pu.repo.GetPermissionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, permission.ErrPermissionNotFound
	}

	return result, nil
}

// SearchPermission implements permission.Service.
func (pu *PermissionUseCase) SearchPermission(ctx context.Context, query *permission.SearchPermissionQuery) (*permission.SearchPermissionResult, error) {
	if query.Page <= 0 {
		query.Page = pu.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = pu.cfg.Pagination.PageLimit
	}

	result, err := pu.repo.SearchPermission(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}

// UpdatePermission implements permission.Service.
func (pu *PermissionUseCase) UpdatePermission(ctx context.Context, cmd *permission.UpdatePermissionCommand) error {
	return pu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := pu.repo.PermissionTaken(ctx, cmd.ID, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			return permission.ErrPermissionNotFound
		}

		if len(result) > 1 || (len(result) == 1 && result[0].ID != cmd.ID) {
			return permission.ErrPermissionAlreadyExists
		}

		err = pu.repo.UpdatePermission(ctx, &permission.UpdatePermissionCommand{
			ID:          cmd.ID,
			Name:        cmd.Name,
			Description: cmd.Description,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
