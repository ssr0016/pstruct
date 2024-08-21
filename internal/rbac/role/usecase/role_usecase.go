package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	rabc "task-management-system/internal/rbac/role"
	"task-management-system/internal/rbac/role/repository/postgres"

	"go.uber.org/zap"
)

type RoleUseCase struct {
	repo *postgres.RoleRository
	cfg  *config.Config
	db   db.DB
	log  *zap.Logger
}

func NewRoleUseCase(db db.DB, cfg *config.Config) rabc.Service {
	return &RoleUseCase{
		repo: postgres.NewRoleRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("role.usecase"),
	}
}

func (r *RoleUseCase) CreateRole(ctx context.Context, cmd *rabc.CreateRoleCommand) error {
	return r.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		err := r.repo.CreateRole(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleUseCase) GetRoleByID(ctx context.Context, id int) (*rabc.RoleDTO, error) {
	result, err := r.repo.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, rabc.ErrRoleNotFound
	}

	return result, nil
}

func (r *RoleUseCase) DeleteRole(ctx context.Context, id int) error {
	return r.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := r.repo.GetRoleByID(ctx, id)
		if err != nil {
			return err
		}

		if result == nil {
			return rabc.ErrRoleNotFound
		}

		err = r.repo.DeleteRole(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleUseCase) GetRoles(ctx context.Context) ([]*rabc.RoleDTO, error) {
	result, err := r.repo.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
