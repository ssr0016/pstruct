package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/department"
	"task-management-system/internal/department/repository/postgres"

	"go.uber.org/zap"
)

type DepartmentUsecase struct {
	repo *postgres.DepartmentRepository
	cfg  *config.Config
	db   db.DB
	log  *zap.Logger
}

func NewDepartmentUsecase(db db.DB, cfg *config.Config) department.Service {
	return &DepartmentUsecase{
		repo: postgres.NewDepartmentRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("department.usecase"),
	}
}

func (du *DepartmentUsecase) CreateDepartment(ctx context.Context, cmd *department.CreateDepartmentCommand) error {
	return du.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := du.repo.DepartmentTaken(ctx, 0, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return department.ErrDepartmentAlreadyExists
		}

		err = du.repo.CreateDepartment(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (du *DepartmentUsecase) GetDepartmentByID(ctx context.Context, id int) (*department.Department, error) {
	result, err := du.repo.GetDepartmentByID(ctx, id)

	if result == nil {
		return nil, department.ErrDepartmentNotFound
	}

	return result, err
}

func (du *DepartmentUsecase) UpdateDepartment(ctx context.Context, cmd *department.UpdateDepartmentCommand) error {
	return du.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := du.repo.DepartmentTaken(ctx, cmd.ID, cmd.Name)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			return department.ErrDepartmentNotFound
		}

		if len(result) > 1 || (len(result) == 1 && result[0].ID != cmd.ID) {
			return department.ErrDepartmentAlreadyExists
		}

		err = du.repo.UpdateDepartment(ctx, &department.UpdateDepartmentCommand{
			ID:   cmd.ID,
			Name: cmd.Name,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (du *DepartmentUsecase) DeleteDepartment(ctx context.Context, id int) error {
	return du.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := du.repo.GetDepartmentByID(ctx, id)
		if err != nil {
			return err
		}

		if result == nil {
			return department.ErrDepartmentNotFound
		}

		err = du.repo.DeleteDepartment(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (du *DepartmentUsecase) SearchDepartment(ctx context.Context, query *department.SearchDepartmentQuery) (*department.SearchDepartmentResult, error) {
	if query.Page <= 0 {
		query.Page = du.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = du.cfg.Pagination.PageLimit
	}

	result, err := du.repo.SearchDepartment(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}
