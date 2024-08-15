package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/task"
	"task-management-system/internal/task/repository/postgres"

	"go.uber.org/zap"
)

type TaskUseCase struct {
	repo *postgres.TaskRepository
	cfg  *config.Config
	db   db.DB
	log  *zap.Logger
}

func NewTaskUsecase(db db.DB, cfg *config.Config) task.Service {
	return &TaskUseCase{
		repo: postgres.NewUserRepository(db), // Ensure this is correct
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("task.usecase"),
	}
}

func (tu *TaskUseCase) CreateTask(ctx context.Context, cmd *task.CreateTaskCommand) error {
	return tu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := tu.repo.TaskTaken(ctx, 0, cmd.Title)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return task.ErrTaskAlreadyExists
		}

		err = tu.repo.Create(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (tu *TaskUseCase) GetTaskByID(ctx context.Context, id int) (*task.Task, error) {
	result, err := tu.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, task.ErrTaskNotFound
	}

	return result, nil
}

func (tu *TaskUseCase) UpdateTask(ctx context.Context, cmd *task.UpdateTaskCommand) error {
	return tu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := tu.repo.TaskTaken(ctx, cmd.ID, cmd.Title)
		if err != nil {
			return err
		}

		if len(result) == 0 {
			return task.ErrTaskNotFound
		}

		if len(result) > 1 || (len(result) == 1 && result[0].ID != cmd.ID) {
			return task.ErrTaskAlreadyExists
		}

		err = tu.repo.Update(ctx, &task.UpdateTaskCommand{
			ID:          cmd.ID,
			Title:       cmd.Title,
			Description: cmd.Description,
			Status:      cmd.Status,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (tu *TaskUseCase) DeleteTask(ctx context.Context, id int) error {
	return tu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := tu.repo.GetByID(ctx, id)
		if err != nil {
			return err
		}

		if result == nil {
			return task.ErrTaskNotFound
		}

		err = tu.repo.Delete(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})

}

func (tu *TaskUseCase) SearchTask(ctx context.Context, query *task.SearchTaskQuery) (*task.SearchTaskResult, error) {
	if query.Page <= 0 {
		query.Page = tu.cfg.Pagination.Page
	}

	if query.PerPage <= 0 {
		query.PerPage = tu.cfg.Pagination.PageLimit
	}

	result, err := tu.repo.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	result.PerPage = query.PerPage
	result.Page = query.Page

	return result, nil
}
