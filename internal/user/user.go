package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, cmd *CreateUserRequest) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
}
