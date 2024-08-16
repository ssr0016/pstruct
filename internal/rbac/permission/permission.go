package permission

import "context"

type Service interface {
	CreatePermission(ctx context.Context, cmd *CreatePermissionCommand) error
	GetPermissionByID(ctx context.Context, id int) (*Permission, error)
	UpdatePermission(ctx context.Context, cmd *UpdatePermissionCommand) error
	DeletePermission(ctx context.Context, id int) error
	SearchPermission(ctx context.Context, query *SearchPermissionQuery) (*SearchPermissionResult, error)
}
