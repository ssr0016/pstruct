package pagination

type Pagination struct {
	Page     int
	PageSize int
	Offset   int
	Limit    int
}

func NewPagination(page, pageSize int) *Pagination {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	limit := pageSize

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
		Limit:    limit,
	}
}
