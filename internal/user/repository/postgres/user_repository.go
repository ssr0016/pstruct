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
				password_hash
			) VALUES (
				$1,
				$2,
				$3,
				$4
			) RETURNING id
		`

		var id int
		err := tx.QueryRow(ctx, rawSQL, cmd.FirstName, cmd.LastName, cmd.Email, cmd.Password).Scan(&id)
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
			email
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
			email
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
