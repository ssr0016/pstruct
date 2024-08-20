package permissions

import (
	"context"
)

type Service interface {
	CreatePermissions(ctx context.Context, cmd *CreatePermissionCommand) error
	GetPermissions(ctx context.Context) ([]*PermissionDTO, error)
	GetPermissionByID(ctx context.Context, id int) (*PermissionDTO, error)
}
