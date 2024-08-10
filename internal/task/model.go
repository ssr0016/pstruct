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

type SearchTaskQuery struct {
	Title       string `schema:"title"`
	Description string `schema:"description"`
	Status      string `schema:"status"`
	PerPage     int    `schema:"per_page"`
	Page        int    `schema:"page"`
}

type SearchTaskResult struct {
	TotalCount int     `json:"total_count"`
	Tasks      []*Task `json:"results"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}
