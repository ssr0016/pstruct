package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

func (u *UserRepository) SearchUser(ctx context.Context, query *user.SearchUserQuery) (*user.SearchUserResult, error) {
	var (
		result = &user.SearchUserResult{
			Users: make([]*user.User, 0),
		}
		sql             bytes.Buffer
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
		paramIndex      = 1
	)

	sql.WriteString(`
		SELECT
			id,
			first_name,
			last_name,
			email,
			address,
			phone_number,
			date_of_birth
		FROM
			users
	`)

	if len(query.FirstName) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("first_name ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.FirstName+"%")
		paramIndex++
	}

	if len(query.LastName) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("last_name ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.LastName+"%")
		paramIndex++
	}

	if len(query.Email) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("email ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Email+"%")
		paramIndex++
	}

	if len(query.Address) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("address ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.Address+"%")
		paramIndex++
	}

	if len(query.PhoneNumber) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("phone_number ILIKE $%d", paramIndex))
		whereParams = append(whereParams, "%"+query.PhoneNumber+"%")
		paramIndex++
	}

	if len(query.DateOfBirth) > 0 {
		whereConditions = append(whereConditions, fmt.Sprintf("date_of_birth = $%d", paramIndex))
		whereParams = append(whereParams, query.DateOfBirth)
		paramIndex++
	}

	if len(whereConditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(whereConditions, " AND "))
	}

	// Getting the count of total results
	count, err := u.getCount(ctx, sql, whereParams)
	if err != nil {
		return nil, err
	}

	if query.PerPage > 0 {
		offset := query.PerPage * (query.Page - 1)
		sql.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramIndex, paramIndex+1))
		whereParams = append(whereParams, query.PerPage, offset)
	}

	err = u.db.Select(ctx, &result.Users, sql.String(), whereParams...)
	if err != nil {
		return nil, err
	}

	result.TotalCount = count

	return result, nil
}

func (r *UserRepository) getCount(ctx context.Context, sql bytes.Buffer, whereParams []interface{}) (int, error) {
	var count int

	rawSQL := "SELECT COUNT(*) FROM (" + sql.String() + ") as t1"

	err := r.db.Get(ctx, &count, rawSQL, whereParams...)
	if err != nil {
		return 0, err
	}

	return count, nil
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
