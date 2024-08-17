package permission

import "context"

type Service interface {
	AddPermission(ctx context.Context, cmd *AddPermissionCommand) error
	GetActionByUserID(ctx context.Context, userID int) ([]*string, error)
	GetListPermission(ctx context.Context, userID int) ([]*ActionScopesPermission, error)
}
