package userroles

import "context"

type Service interface {
	AssignRoleToUser(ctx context.Context, cmd *CreateUserRoleCommand) error
	RemoveRoleFromUser(ctx context.Context, cmd *CreateRemoveUserRoleCommand) error
	SearchUserRole(ctx context.Context, query *SearchUserRoleQuery) (*SearchUserRoleResult, error)
}
