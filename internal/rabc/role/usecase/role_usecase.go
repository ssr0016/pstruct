package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	rabc "task-management-system/internal/rabc/role"
	"task-management-system/internal/rabc/role/repository/postgres"

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
		result, err := r.repo.RoleTaken(ctx, 0, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return rabc.ErrRoleAlreadyExists
		}

		err = r.repo.CreateRole(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *RoleUseCase) GetRoleByID(ctx context.Context, id int) (*rabc.Role, error) {
	result, err := r.repo.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, rabc.ErrRoleNotFound
	}

	return result, nil
}

func (r *RoleUseCase) UpdateRole(ctx context.Context, cmd *rabc.UpdateRoleCommand) error {
	return r.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := r.repo.RoleTaken(ctx, cmd.ID, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			return rabc.ErrRoleNotFound
		}

		if len(result) > 1 || (len(result) == 1 && result[0].ID != cmd.ID) {
			return rabc.ErrRoleAlreadyExists
		}

		err = r.repo.UpdateRole(ctx, &rabc.UpdateRoleCommand{
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

func (r *RoleUseCase) SearchRole(ctx context.Context, query *rabc.SearchRoleQuery) (*rabc.SearchRoleResult, error) {
	if query.Page <= 0 {
		query.Page = r.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = r.cfg.Pagination.PageLimit
	}

	result, err := r.repo.SearchRole(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}
