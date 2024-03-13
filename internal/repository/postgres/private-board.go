package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/4aykovski/task-manager-api/internal/model"
	"github.com/4aykovski/task-manager-api/internal/repository"
	"github.com/4aykovski/task-manager-api/pkg/database/postgres"
	"github.com/lib/pq"
)

type PrivateBoardRepository struct {
	db *postgres.Postgres
}

func NewPrivateBoardRepository(db *postgres.Postgres) *PrivateBoardRepository {
	return &PrivateBoardRepository{db: db}
}

func (repo *PrivateBoardRepository) CreatePrivateBoard(ctx context.Context, board *model.PrivateBoard) error {
	const op = "internal.repository.postgres.privateboard.InsertPrivateBoard"
	const insertPrivateBoardStmt = `INSERT INTO private_boards (name, color, user_id) 
					VALUES ($1, $2, $3)`

	stmt, err := repo.db.PrepareContext(ctx, insertPrivateBoardStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, board.Name, board.Color, board.UserId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrPrivateBoardAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PrivateBoardRepository) GetPrivateBoards(ctx context.Context, userId string) ([]model.PrivateBoard, error) {
	const op = "internal.repository.postgres.privateboard.GetPrivateBoards"
	const getPrivateBoardsStmt = `SELECT id, name, color, user_id FROM private_boards WHERE user_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getPrivateBoardsStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var privateBoards []model.PrivateBoard
	for rows.Next() {
		var privateBoard model.PrivateBoard
		err = rows.Scan(
			&privateBoard.Id,
			&privateBoard.Name,
			&privateBoard.Color,
			&privateBoard.UserId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		privateBoards = append(privateBoards, privateBoard)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(privateBoards) == 0 {
		return nil, repository.ErrPrivateBoardNotFound
	}

	return privateBoards, nil
}

func (repo *PrivateBoardRepository) DeletePrivateBoard(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.privateboard.DeletePrivateBoard"
	const deletePrivateBoardStmt = `DELETE FROM private_boards WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deletePrivateBoardStmt)
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
		return repository.ErrPrivateBoardNotFound
	}

	return nil
}

func (repo *PrivateBoardRepository) UpdatePrivateBoard(ctx context.Context, board *model.PrivateBoard) error {
	const op = "internal.repository.postgres.privateboard.UpdatePrivateBoard"
	const updatePrivateBoardStmt = `UPDATE private_boards SET name = $1, color = $2 WHERE id = $3`

	stmt, err := repo.db.PrepareContext(ctx, updatePrivateBoardStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, board.Name, board.Color, board.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrPrivateBoardNotFound
	}

	return nil
}
