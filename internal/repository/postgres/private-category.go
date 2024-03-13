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

type PrivateCategoryRepository struct {
	db *postgres.Postgres
}

func NewPrivateCategoryRepository(db *postgres.Postgres) *PrivateCategoryRepository {
	return &PrivateCategoryRepository{db: db}
}

func (repo *PrivateCategoryRepository) CreatePrivateCategory(ctx context.Context, category *model.PrivateCategory) error {
	const op = "internal.repository.postgres.privateCategory.InsertPrivateCategory"
	const insertPrivateCategoryStmt = `INSERT INTO private_categories (name, color, board_id)
					VALUES ($1, $2, $3)`

	stmt, err := repo.db.PrepareContext(ctx, insertPrivateCategoryStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, category.Name, category.Color, category.BoardId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrPrivateCategoryAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PrivateCategoryRepository) GetPrivateCategories(ctx context.Context, boardId int) ([]model.PrivateCategory, error) {
	const op = "internal.repository.postgres.privateCategory.GetPrivateCategories"
	const getPrivateCategoriesStmt = `SELECT id, name, color, board_id FROM private_categories WHERE board_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getPrivateCategoriesStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var categories []model.PrivateCategory
	for rows.Next() {
		var category model.PrivateCategory
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
		return nil, repository.ErrPrivateCategoryNotFound
	}

	return categories, nil
}

func (repo *PrivateCategoryRepository) DeletePrivateCategory(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.privateCategory.DeletePrivateCategory"
	const deletePrivateCategoryStmt = `DELETE FROM private_categories WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deletePrivateCategoryStmt)
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
		return repository.ErrPrivateCategoryNotFound
	}

	return nil
}

func (repo *PrivateCategoryRepository) UpdatePrivateCategory(ctx context.Context, category *model.PrivateCategory) error {
	const op = "internal.repository.postgres.privateCategory.UpdatePrivateCategory"
	const updatePrivateCategoryStmt = `UPDATE private_categories SET name = $1, color = $2 WHERE id = $3`

	stmt, err := repo.db.PrepareContext(ctx, updatePrivateCategoryStmt)
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
		return repository.ErrPrivateCategoryNotFound
	}

	return nil
}
