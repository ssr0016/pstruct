package permissionuser

import "context"

type Service interface {
	CreateUserPermission(ctx context.Context, cmd *CreateUserPermissionCommand) error
	GetUsersPermissions(ctx context.Context, query *UserPermissionsQuery) (*UserPermissionsResult, error)
	GetUserPermissionByID(ctx context.Context, id int) (*UserPermission, error)
	DeleteUserPermission(ctx context.Context, id int) error
	GetAllUserPermissions(ctx context.Context, userID string) ([]*UserPermission, error)
}
