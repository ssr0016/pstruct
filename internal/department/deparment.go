package department

import "context"

type Service interface {
	CreateDepartment(ctx context.Context, cmd *CreateDepartmentCommand) error
	GetDepartmentByID(ctx context.Context, id int) (*Department, error)
	UpdateDepartment(ctx context.Context, cmd *UpdateDepartmentCommand) error
	DeleteDepartment(ctx context.Context, id int) error
	SearchDepartment(ctx context.Context, query *SearchDepartmentQuery) (*SearchDepartmentResult, error)
}
