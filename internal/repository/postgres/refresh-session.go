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

type RefreshSessionRepository struct {
	db *postgres.Postgres
}

func NewRefreshSessionRepository(db *postgres.Postgres) *RefreshSessionRepository {
	return &RefreshSessionRepository{db}
}

func (repo *RefreshSessionRepository) CreateRefreshSession(ctx context.Context, refreshSession *model.RefreshSession) error {
	const op = "internal.repository.postgres.refreshSession.CreateRefreshSession"
	const createRefreshSessionStmt = `INSERT INTO refresh_sessions (token, user_id, expires_in, fingerprint) 
										VALUES ($1, $2, $3, $4)`

	stmt, err := repo.db.PrepareContext(ctx, createRefreshSessionStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, refreshSession.Token, refreshSession.UserId, refreshSession.ExpiresIn, refreshSession.Fingerprint)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrRefreshSessionAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *RefreshSessionRepository) GetRefreshSessionWithToken(ctx context.Context, token string) (*model.RefreshSession, error) {
	const op = "internal.repository.postgres.refreshSession.GetRefreshSessionWithToken"
	const getRefreshSessionStmt = `SELECT token, user_id, expires_in, fingerprint 
									FROM refresh_sessions WHERE token = $1`

	stmt, err := repo.db.PrepareContext(ctx, getRefreshSessionStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var refreshSession model.RefreshSession
	err = stmt.QueryRowContext(ctx, token).Scan(
		&refreshSession.Token,
		&refreshSession.UserId,
		&refreshSession.ExpiresIn,
		&refreshSession.Fingerprint,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrRefreshSessionNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &refreshSession, nil
}

func (repo *RefreshSessionRepository) DeleteRefreshSession(ctx context.Context, token string) error {
	const op = "internal.repository.postgres.refreshSession.DeleteRefreshSession"
	const deleteRefreshSessionStmt = `DELETE FROM refresh_sessions WHERE token = $1`

	stmt, err := repo.db.PrepareContext(ctx, deleteRefreshSessionStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, token)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrRefreshSessionNotFound
	}

	return nil
}
