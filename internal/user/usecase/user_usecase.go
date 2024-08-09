package usecase

import (
	"task-management-system/internal/user"
	"task-management-system/internal/user/repository/postgres"
)

type UserUsecase struct {
	Repo *postgres.UserRepository
}

func (uc *UserUsecase) CreateUser(u *user.User) error {
	return uc.Repo.Create(u)
}

func (uc *UserUsecase) GetUserByID(id int) (*user.User, error) {
	return uc.Repo.GetByID(id)
}

func (uc *UserUsecase) UpdateUser(u *user.User) error {
	return uc.Repo.Update(u)
}

func (uc *UserUsecase) DeleteUser(id int) error {
	return uc.Repo.Delete(id)
}

func (uc *UserUsecase) GetAllUsers() ([]user.User, error) {
	return uc.Repo.GetAll()
}
