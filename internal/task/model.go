package task

type Task struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Status      string `db:"status" json:"status"`
}

type CreateTaskCommand struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTaskCommand struct {
	ID          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type SearchTaskQuery struct{}

type SearchTaskResult struct{}
