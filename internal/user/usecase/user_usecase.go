package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/user"
	"task-management-system/internal/user/repository/postgres"
	util "task-management-system/pkg/util/password"

	"go.uber.org/zap"
)

type UserUsecase struct {
	repo *postgres.UserRepository
	cfg  *config.Config
	log  *zap.Logger
	db   db.DB
}

func NewUserCase(db db.DB, cfg *config.Config) user.Service {
	return &UserUsecase{
		repo: postgres.NewUserRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("user.usecase"),
	}

}

func (uu *UserUsecase) CreateUser(ctx context.Context, cmd *user.CreateUserRequest) error {
	return uu.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		result, err := uu.repo.TaskTaken(ctx, 0, cmd.Email)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			return user.ErrUserAlreadyExists
		}

		// Hash the password
		passwordHash, err := util.HashPassword(cmd.Password)
		if err != nil {
			return err
		}

		cmd.Password = passwordHash

		err = uu.repo.Create(ctx, cmd)
		if err != nil {
			return err
		}

		return nil
	})
}

func (uu *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	result, err := uu.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uu *UserUsecase) GetUserByID(ctx context.Context, id int) (*user.User, error) {
	result, err := uu.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
