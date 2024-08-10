package config

import (
	"os"
	"strconv"
	"task-management-system/internal/pagination"
)

func (cfg *Config) LoadPaginationConfig() {
	page, err := strconv.Atoi(os.Getenv("PAGE"))
	if err != nil {
		page = pagination.DefaultPage
	}
	cfg.Pagination.Page = page

	pageLimit, err := strconv.Atoi(os.Getenv("PER_PAGE"))
	if err != nil {
		pageLimit = pagination.DefaultPageLimit
	}
	cfg.Pagination.PageLimit = pageLimit
}
