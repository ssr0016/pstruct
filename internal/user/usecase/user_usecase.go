package usecase

import (
	"task-management-system/internal/db"
	"task-management-system/internal/user/repository/postgres"
)

type UserUsecase struct {
	repo *postgres.UserRepository
}

func NewUserCase(db db.DB) *UserUsecase {
	return &UserUsecase{
		repo: postgres.NewUserRepository(db),
	}

}

// func (s *UserUsecase) CreateUser(u *user.User) error {
// 	return s.repo.Create(u)
// }

// func (s *UserUsecase) GetUserByID(id int) (*user.User, error) {
// 	return s.repo.GetByID(id)
// }

// func (s *UserUsecase) UpdateUser(u *user.User) error {
// 	return s.repo.Update(u)
// }

// func (s *UserUsecase) DeleteUser(id int) error {
// 	return s.repo.Delete(id)
// }

// func (s *UserUsecase) GetAllUsers() ([]user.User, error) {
// 	return s.repo.GetAll()
// }
