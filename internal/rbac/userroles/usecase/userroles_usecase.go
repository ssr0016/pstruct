package usecase

import (
	"context"
	"task-management-system/config"
	"task-management-system/internal/db"
	"task-management-system/internal/rbac/userroles"
	"task-management-system/internal/rbac/userroles/repository/postgres"

	"go.uber.org/zap"
)

type UserRoleUseCase struct {
	repo *postgres.UserRoleRepository
	db   db.DB
	cfg  *config.Config
	log  *zap.Logger
}

func NewUserRoleUseCase(db db.DB, cfg *config.Config) userroles.Service {
	return &UserRoleUseCase{
		repo: postgres.NewUserRoleRepository(db),
		db:   db,
		cfg:  cfg,
		log:  zap.L().Named("userroles.usecase"),
	}
}

func (uru *UserRoleUseCase) Assign(ctx context.Context, userID, roleID int) error {
	if userID <= 0 {
		return userroles.ErrInvalidUserID
	}

	if roleID <= 0 {
		return userroles.ErrInvalidRoleID
	}

	err := uru.repo.Assign(ctx, userID, roleID)
	if err != nil {
		uru.log.Error("error assigning user role", zap.Error(err))
		return err
	}

	return nil

}
