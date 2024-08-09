package postgres

import (
	"task-management-system/internal/task"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	DB *sqlx.DB
}

func (r *TaskRepository) Create(task *task.Task) error {
	_, err := r.DB.NamedExec(`INSERT INTO tasks (title, description, status) VALUES (:title, :description, :status)`, task)
	return err
}

func (r *TaskRepository) GetByID(id int) (*task.Task, error) {
	var t task.Task
	err := r.DB.Get(&t, "SELECT * FROM tasks WHERE id=$1", id)
	return &t, err
}

func (r *TaskRepository) Update(task *task.Task) error {
	_, err := r.DB.NamedExec(`UPDATE tasks SET title=:title, description=:description, status=:status WHERE id=:id`, task)
	return err
}

func (r *TaskRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	return err
}

func (r *TaskRepository) GetAll() ([]task.Task, error) {
	var tasks []task.Task
	err := r.DB.Select(&tasks, "SELECT * FROM tasks")
	return tasks, err
}
