package usecase

import (
	"context"
	"errors"
	"task-management-system/internal/task"
	"task-management-system/internal/task/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type TaskUseCase struct {
	repo *postgres.TaskRepository
}

func NewTaskUsecase(db *sqlx.DB) task.Service {
	return &TaskUseCase{
		repo: postgres.NewUserRepository(db),
	}
}

func (tu *TaskUseCase) CreateTask(ctx context.Context, cmd *task.CreateTaskCommand) error {
	return tu.repo.Create(ctx, cmd)
}

func (tu *TaskUseCase) GetTaskByID(ctx context.Context, id int) (*task.Task, error) {
	result, err := tu.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tu *TaskUseCase) UpdateTask(ctx context.Context, cmd *task.UpdateTaskCommand) error {
	// Fetch the existing task to ensure it exists
	existingTask, err := tu.repo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	// Check if task exists
	if existingTask == nil {
		return errors.New("task not found")
	}

	// Update the task in the repository
	return tu.repo.Update(ctx, cmd)
}

func (tu *TaskUseCase) DeleteTask(ctx context.Context, id int) error {
	existingTask, err := tu.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if existingTask == nil {
		return errors.New("task not found")
	}

	err = tu.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (tu *TaskUseCase) SearchTask(ctx context.Context, query *task.SearchTaskQuery) (*task.SearchTaskResult, error) {
	panic("unimplemented")
}
