package usecase

import (
	"task-management-system/internal/task"
	"task-management-system/internal/task/repository/postgres"
)

type TaskUsecase struct {
	Repo *postgres.TaskRepository
}

func (uc *TaskUsecase) CreateTask(t *task.Task) error {
	return uc.Repo.Create(t)
}

func (uc *TaskUsecase) GetTaskByID(id int) (*task.Task, error) {
	return uc.Repo.GetByID(id)
}

func (uc *TaskUsecase) UpdateTask(t *task.Task) error {
	return uc.Repo.Update(t)
}

func (uc *TaskUsecase) DeleteTask(id int) error {
	return uc.Repo.Delete(id)
}

func (uc *TaskUsecase) GetAllTasks() ([]task.Task, error) {
	return uc.Repo.GetAll()
}
