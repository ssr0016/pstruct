package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, cmd *CreateUserRequest) error
	GetUserByEmail(ctx context.Context, cmd *LoginUserRequest) (string, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, cmd *UpdateUserRequest) error
	DeleteUser(ctx context.Context, id int) error
	SearchUser(ctx context.Context, query *SearchUserQuery) (*SearchUserResult, error)
}
