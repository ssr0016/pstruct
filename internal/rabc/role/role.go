package rabc

import (
	"context"
)

type Service interface {
	CreateRole(ctx context.Context, cmd *CreateRoleCommand) error
	GetRoleByID(ctx context.Context, id int) (*Role, error)
	UpdateRole(ctx context.Context, cmd *UpdateRoleCommand) error
	DeleteRole(ctx context.Context, id int) error
	SearchRole(ctx context.Context, query *SearchRoleQuery) (*SearchRoleResult, error)
}
