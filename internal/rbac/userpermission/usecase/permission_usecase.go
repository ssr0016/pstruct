package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	permission "task-management-system/internal/rbac/userpermission"
	"task-management-system/internal/rbac/userpermission/reporitory/postgres"

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

func (p *PermissionUseCase) AddPermission(ctx context.Context, cmd *permission.AddPermissionCommand) error {
	return p.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		err := p.repo.Delete(ctx, cmd.UserID)
		if err != nil {
			return err
		}

		if len(cmd.Permission) == 0 {
			return nil
		}

		if len(cmd.Permission) > 0 {
			var permissions []*permission.UserPermission
			for _, p := range cmd.Permission {
				permissions = append(permissions, &permission.UserPermission{
					UserID: cmd.UserID,
					Action: p.Action,
					Scope:  p.Scope,
				})
			}

			err := p.repo.Add(ctx, permissions)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (p *PermissionUseCase) GetActionByUserID(ctx context.Context, userID int) ([]*string, error) {
	result, err := p.repo.GetActionByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PermissionUseCase) GetListPermission(ctx context.Context, userID int) ([]*permission.ActionScopesPermission, error) {
	result, err := p.repo.GetListPermission(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
