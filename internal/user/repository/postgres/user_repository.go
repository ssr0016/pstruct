package postgres

import (
	"task-management-system/internal/db"
)

type UserRepository struct {
	DB db.DB
}

func NewUserRepository(db db.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// func (r *UserRepository) Create(user *user.User) error {
// 	_, err := r.DB.NamedExec(`INSERT INTO users (username, password) VALUES (:username, :password)`, user)
// 	return err
// }

// func (r *UserRepository) GetByID(id int) (*user.User, error) {
// 	var u user.User
// 	err := r.DB.Get(&u, "SELECT * FROM users WHERE id=$1", id)
// 	return &u, err
// }

// func (r *UserRepository) Update(user *user.User) error {
// 	_, err := r.DB.NamedExec(`UPDATE users SET username=:username, password=:password WHERE id=:id`, user)
// 	return err
// }

// func (r *UserRepository) Delete(id int) error {
// 	_, err := r.DB.Exec("DELETE FROM users WHERE id=$1", id)
// 	return err
// }

// func (r *UserRepository) GetAll() ([]user.User, error) {
// 	var users []user.User
// 	err := r.DB.Select(&users, "SELECT * FROM users")
// 	return users, err
// }
