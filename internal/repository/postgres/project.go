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

type ProjectRepository struct {
	db *postgres.Postgres
}

func NewProjectRepository(db *postgres.Postgres) *ProjectRepository {
	return &ProjectRepository{db}
}

func (repo *ProjectRepository) CreateProject(ctx context.Context, project *model.Project) error {
	const op = "internal.repository.postgres.project.CreateProject"
	const createProjectStmt = `INSERT INTO projects (name, description, owner) 
					VALUES ($1, $2, $3)`

	stmt, err := repo.db.PrepareContext(ctx, createProjectStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, project.Name, project.Description, project.Owner)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrProjectAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *ProjectRepository) GetProjectWithId(ctx context.Context, id int) (*model.Project, error) {
	const op = "internal.repository.postgres.project.GetProject"
	const createProjectStmt = `SELECT id, name, description, owner
								FROM projects WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, createProjectStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var project model.Project
	err = stmt.QueryRowContext(ctx, id).Scan(
		&project.Id,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.Owner,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrProjectNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &project, nil
}

func (repo *ProjectRepository) GetProjects(ctx context.Context) ([]model.Project, error) {
	const op = "internal.repository.postgres.project.GetProjects"
	const createProjectStmt = `SELECT id, name, description, owner
								FROM projects`

	stmt, err := repo.db.PrepareContext(ctx, createProjectStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var project model.Project
		err = rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.CreatedAt,
			&project.Owner,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(projects) == 0 {
		return nil, repository.ErrProjectNotFound
	}

	return projects, nil
}

func (repo *ProjectRepository) UpdateProject(ctx context.Context, project *model.Project) error {
	const op = "internal.repository.postgres.project.UpdateProject"
	const createProjectStmt = `UPDATE projects SET name = $1, description = $2
								WHERE id = $3`

	stmt, err := repo.db.PrepareContext(ctx, createProjectStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, project.Name, project.Description, project.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *ProjectRepository) DeleteProject(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.project.DeleteProject"
	const createProjectStmt = `DELETE FROM projects WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, createProjectStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
