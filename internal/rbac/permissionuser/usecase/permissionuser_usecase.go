package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"

	"task-management-system/internal/rbac/permissionuser"
	"task-management-system/internal/rbac/permissionuser/repository/postgres"

	"go.uber.org/zap"
)

type PermissionUserUseCase struct {
	repo *postgres.PermissionUserRepository
	db   db.DB
	cfg  *config.Config
	log  *zap.Logger
}

func NewPermissionUserUseCase(db db.DB, cfg *config.Config) permissionuser.Service {
	return &PermissionUserUseCase{
		repo: postgres.NewPermissionUserRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("permissionuser.usecase"),
	}
}

// CreateUserPermission implements permissionuser.Service.
func (puu *PermissionUserUseCase) CreateUserPermission(ctx context.Context, cmd *permissionuser.CreateUserPermissionCommand) error {
	return puu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		err := puu.repo.CreateUserPermission(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

// DeleteUserPermission implements permissionuser.Service.
func (puu *PermissionUserUseCase) DeleteUserPermission(ctx context.Context, id int) error {
	return puu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := puu.repo.GetUserPermissionByID(ctx, id)
		if err != nil {
			return err
		}

		if result == nil {
			return permissionuser.ErrUserPermissionNotFound
		}

		err = puu.repo.DeleteUserPermission(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

// GetUsersPermissions implements permissionuser.Service.
func (puu *PermissionUserUseCase) GetUsersPermissions(ctx context.Context, query *permissionuser.UserPermissionsQuery) (*permissionuser.UserPermissionsResult, error) {
	if query.Page <= 0 {
		query.Page = puu.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = puu.cfg.Pagination.PageLimit
	}

	result, err := puu.repo.GetUserPermissions(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}

func (p *PermissionUserUseCase) GetUserPermissionByID(ctx context.Context, id int) (*permissionuser.UserPermission, error) {
	result, err := p.repo.GetUserPermissionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PermissionUserUseCase) GetAllUserPermissions(ctx context.Context, userID string) ([]*permissionuser.UserPermission, error) {
	result, err := p.repo.GetAllUserPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
