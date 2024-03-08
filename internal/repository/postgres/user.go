package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/4aykovski/task-manager-api/internal/model"
	"github.com/4aykovski/task-manager-api/internal/repository"
	"github.com/4aykovski/task-manager-api/pkg/database/postgres"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	const op = "internal.repository.postgres.user.CreateUser"
	const createUserStmt = `INSERT INTO users (login, email, password, language) 
					VALUES ($1, $2, $3, $4)`

	stmt, err := repo.db.PrepareContext(ctx, createUserStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Login, user.Email, user.Password, user.Language)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrUserAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *UserRepository) GetUserWithId(ctx context.Context, id int) (*model.User, error) {
	const op = "internal.repository.postgres.user.GetUser"
	const getUserStmt = `SELECT id, login, email, password, registered_at, last_seen_at, is_admin, language, theme, about
					FROM users WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getUserStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var user model.User
	err = stmt.QueryRowContext(ctx, id).Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Password,
		&user.RegisteredAt,
		&user.LastSeenAt,
		&user.IsAdmin,
		&user.Language,
		&user.Theme,
		&user.About,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (repo *UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	const op = "internal.repository.postgres.user.GetUsers"
	const getUsersStmt = `SELECT id, login, email, password, registered_at, last_seen_at, is_admin, language, theme, about
					FROM users`

	stmt, err := repo.db.PrepareContext(ctx, getUsersStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Email,
			&user.Password,
			&user.RegisteredAt,
			&user.LastSeenAt,
			&user.IsAdmin,
			&user.Language,
			&user.Theme,
			&user.About,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(users) == 0 {
		return nil, repository.ErrUserNotFound
	}

	return nil, err

}

func (repo *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	const op = "internal.repository.postgres.user.GetUser"
	const updateUserStmt = `UPDATE users SET login = $1, email = $2, password = $3, language = $4, theme = $5, about = $6, last_seen_at = $7, is_admin = $8
					WHERE id = $9`

	stmt, err := repo.db.PrepareContext(ctx, updateUserStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, user.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrUserNotFound
	}

	return nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.user.GetUser"
	const deleteUserStmt = `DELETE FROM users WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deleteUserStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrUserNotFound
	}

	return nil
}
