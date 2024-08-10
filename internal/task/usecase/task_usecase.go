package usecase

import (
	"task-management-system/internal/task"
	"task-management-system/internal/task/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type TaskUseCase struct {
	repo *postgres.TaskRepository
}

func NewTaskUsecase(db *sqlx.DB) *TaskUseCase {
	return &TaskUseCase{
		repo: postgres.NewUserRepository(db),
	}
}

func (tu *TaskUseCase) CreateTask(t *task.Task) error {
	return tu.repo.Create(t)
}

func (tu *TaskUseCase) GetTaskByID(id int) (*task.Task, error) {
	return tu.repo.GetByID(id)
}

func (tu *TaskUseCase) UpdateTask(t *task.Task) error {
	return tu.repo.Update(t)
}

func (tu *TaskUseCase) DeleteTask(id int) error {
	return tu.repo.Delete(id)
}

func (tu *TaskUseCase) GetAllTasks() ([]task.Task, error) {
	return tu.repo.GetAll()
}
