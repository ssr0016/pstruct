package userroles

import "context"

type Service interface {
	AssignRoles(ctx context.Context, cmd *CreateUserRolesCommand) error
	RemoveUserRoles(ctx context.Context, cmd *RemoveUserRolesCommand) error
	GetUserRolesByID(ctx context.Context, userID string) ([]*UserRole, error)
}
