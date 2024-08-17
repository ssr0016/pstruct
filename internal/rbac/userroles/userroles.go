package userroles

import "context"

type Service interface {
	AssignRoles(ctx context.Context, cmd *CreateUserRolesCommand) error
	RemoveUserRoles(ctx context.Context, cmd *RemoveUserRolesCommand) error
	GetUserRolesByID(ctx context.Context, id int) ([]*UserRole, error)
	UpdateUserRoles(ctx context.Context, cmd *UpdateUserRolesCommand) error
	SearchUserRoles(ctx context.Context, query *SearchUserRolesQuery) (*SearchUserRolesResult, error)
}
