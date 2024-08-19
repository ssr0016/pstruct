package permissions

import (
	"context"
)

type Service interface {
	CreatePermissions(ctx context.Context, cmd *CreatePermissionCommand) error
	GetPermissions(ctx context.Context) ([]*Permission, error)
	GetPermissionByID(ctx context.Context, id int) (*Permission, error)
}
