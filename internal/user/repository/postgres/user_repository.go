package postgres

import (
	"context"
	"database/sql"
	"errors"
	"task-management-system/internal/db"
	"task-management-system/internal/user"

	"go.uber.org/zap"
)

type UserRepository struct {
	db     db.DB
	logger *zap.Logger
}

func NewUserRepository(db db.DB) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: zap.L().Named("user.repository"),
	}
}

func (u *UserRepository) Create(ctx context.Context, cmd *user.CreateUserRequest) error {
	return u.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			INSERT INTO users (
				first_name,
				last_name,
				email,
				password_hash, 
			 	address,
				phone_number,
				date_of_birth
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7
			) RETURNING id
		`

		var id int

		err := tx.QueryRow(ctx, rawSQL, cmd.FirstName, cmd.LastName, cmd.Email, cmd.Password, cmd.Address, cmd.PhoneNumber, cmd.DateOfBirth).Scan(&id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var user user.User

	rawSQL := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password_hash,
			address,
			phone_number,
			date_of_birth
		FROM
			users
		WHERE
			email = $1
	`

	err := u.db.Get(ctx, &user, rawSQL, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) GetUserByID(ctx context.Context, id int) (*user.User, error) {
	var user user.User

	rawSQL := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password_hash,
			address,
			phone_number,
			date_of_birth
		FROM
			users
		WHERE
			id = $1
	`

	err := u.db.Get(ctx, &user, rawSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) Update(ctx context.Context, cmd *user.UpdateUserRequest) error {
	return u.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
		UPDATE users
		SET
			first_name = $1,
			last_name = $2,
			email = $3,
			address = $4,
			phone_number = $5,
			date_of_birth = $6
		WHERE
			id = $7
	`

		_, err := u.db.Exec(ctx, rawSQL, cmd.FirstName, cmd.LastName, cmd.Email, cmd.Address, cmd.PhoneNumber, cmd.DateOfBirth, cmd.ID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *UserRepository) Delete(ctx context.Context, id int) error {
	return u.db.WithTransaction(ctx, func(ctx context.Context, tx db.Tx) error {
		rawSQL := `
			DELETE FROM users
			WHERE id = $1
		`
		_, err := tx.Exec(ctx, rawSQL, id)
		return err
	})
}

func (u *UserRepository) TaskTaken(ctx context.Context, id int, email string) ([]*user.User, error) {
	var result []*user.User

	rawSQL := `
		SELECT
			id,
			first_name,
			last_name,
			email
		FROM
			users
		WHERE
			email = $1
	`

	err := u.db.Select(ctx, &result, rawSQL, email)
	if err != nil {
		return nil, err
	}

	return result, nil
}
