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

type ProjectBoardRepository struct {
	db *postgres.Postgres
}

func NewProjectBoardRepository(db *postgres.Postgres) *ProjectBoardRepository {
	return &ProjectBoardRepository{db: db}
}

func (repo *ProjectBoardRepository) CreateProjectBoard(ctx context.Context, board *model.ProjectBoard) error {
	const op = "internal.repository.postgres.projectboard.InsertProjectBoard"
	const createProjectBoardStmt = `INSERT INTO project_boards (name, color, project_id)
					VALUES ($1, $2, $3)`

	stmt, err := repo.db.PrepareContext(ctx, createProjectBoardStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, board.Name, board.Color, board.ProjectId)
	if err != nil {

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrProjectBoardAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *ProjectBoardRepository) GetProjectBoards(ctx context.Context, projectId int) ([]model.ProjectBoard, error) {
	const op = "internal.repository.postgres.projectboard.GetProjectBoards"
	const getProjectBoardsStmt = `SELECT id, name, color, project_id FROM project_boards WHERE project_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectBoardsStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var projectBoards []model.ProjectBoard
	for rows.Next() {
		var projectBoard model.ProjectBoard
		err = rows.Scan(
			&projectBoard.Id,
			&projectBoard.Name,
			&projectBoard.Color,
			&projectBoard.ProjectId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projectBoards = append(projectBoards, projectBoard)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(projectBoards) == 0 {
		return nil, repository.ErrProjectBoardNotFound
	}

	return projectBoards, nil
}

func (repo *ProjectBoardRepository) DeleteProjectBoard(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.projectboard.DeleteProjectBoard"
	const deleteProjectBoardStmt = `DELETE FROM project_boards WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deleteProjectBoardStmt)
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
		return repository.ErrProjectBoardNotFound
	}

	return nil
}

func (repo *ProjectBoardRepository) UpdateProjectBoard(ctx context.Context, board *model.ProjectBoard) error {
	const op = "internal.repository.postgres.projectboard.UpdateProjectBoard"
	const updateProjectBoardStmt = `UPDATE project_boards SET name = $1, color = $2 WHERE id = $3`

	stmt, err := repo.db.PrepareContext(ctx, updateProjectBoardStmt)
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
		return repository.ErrProjectBoardNotFound
	}

	return nil
}
