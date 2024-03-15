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

type ProjectCategoryRepository struct {
	db *postgres.Postgres
}

func NewProjectCategoryRepository(db *postgres.Postgres) *ProjectCategoryRepository {
	return &ProjectCategoryRepository{db: db}
}

func (repo *ProjectCategoryRepository) CreateProjectCategory(ctx context.Context, category *model.ProjectCategory) error {
	const op = "internal.repository.postgres.projectCategory.InsertProjectCategory"
	const createProjectCategoryStmt = `INSERT INTO project_categories (name, color, board_id)
					VALUES ($1, $2, $3)`

	stmt, err := repo.db.PrepareContext(ctx, createProjectCategoryStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, category.Name, category.Color, category.BoardId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrProjectCategoryAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *ProjectCategoryRepository) GetProjectCategories(ctx context.Context, boardId int) ([]model.ProjectCategory, error) {

	const op = "internal.repository.postgres.projectCategory.GetProjectCategories"
	const getProjectCategoriesStmt = `SELECT id, name, color, board_id FROM project_categories WHERE board_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectCategoriesStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var categories []model.ProjectCategory
	for rows.Next() {
		var category model.ProjectCategory
		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.Color,
			&category.BoardId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(categories) == 0 {
		return nil, repository.ErrProjectCategoryNotFound
	}

	return categories, nil
}

func (repo *ProjectCategoryRepository) DeleteProjectCategory(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.projectCategory.DeleteProjectCategory"
	const deleteProjectCategoryStmt = `DELETE FROM project_categories WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deleteProjectCategoryStmt)
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
		return repository.ErrProjectCategoryNotFound
	}

	return nil
}

func (repo *ProjectCategoryRepository) UpdateProjectCategory(ctx context.Context, category *model.ProjectCategory) error {
	const op = "internal.repository.postgres.projectCategory.UpdateProjectCategory"
	const updateProjectCategoryStmt = `UPDATE project_categories SET name = $1, color = $2 WHERE id = $3`

	stmt, err := repo.db.PrepareContext(ctx, updateProjectCategoryStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, category.Name, category.Color, category.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrProjectCategoryNotFound
	}

	return nil
}
