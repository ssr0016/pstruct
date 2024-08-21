package rabc

import (
	"context"
)

type Service interface {
	CreateRole(ctx context.Context, cmd *CreateRoleCommand) error
	GetRoleByID(ctx context.Context, id int) (*RoleDTO, error)
	DeleteRole(ctx context.Context, id int) error
	GetRoles(ctx context.Context) ([]*RoleDTO, error)
}
