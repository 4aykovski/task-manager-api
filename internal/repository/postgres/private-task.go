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

type PrivateTaskRepository struct {
	db *postgres.Postgres
}

func NewPrivateTaskRepository(db *postgres.Postgres) *PrivateTaskRepository {
	return &PrivateTaskRepository{db}
}

func (repo *PrivateTaskRepository) CreatePrivateTask(ctx context.Context, privateTask *model.PrivateTask) error {
	const op = "internal.repository.postgres.privateTask.CreatePrivateTask"
	const createPrivateTaskStmt = `INSERT INTO private_tasks (category_id, name, description, deadline, user_id, board_id)
					VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := repo.db.PrepareContext(ctx, createPrivateTaskStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, privateTask.CategoryId, privateTask.Name, privateTask.Description, privateTask.Deadline, privateTask.UserId, privateTask.BoardId)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrPrivateTaskAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PrivateTaskRepository) GetPrivateTasksWithUserId(ctx context.Context, userId string) ([]model.PrivateTask, error) {
	const op = "internal.repository.postgres.privateTask.GetPrivateTasks"
	const getPrivateTasksStmt = `SELECT id, category_id, name, description, status, date_create, deadline, user_id, board_id
					FROM private_tasks WHERE user_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getPrivateTasksStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var privateTasks []model.PrivateTask
	for rows.Next() {
		var privateTask model.PrivateTask
		err = rows.Scan(
			&privateTask.Id,
			&privateTask.CategoryId,
			&privateTask.Name,
			&privateTask.Description,
			&privateTask.Status,
			&privateTask.DateCreate,
			&privateTask.Deadline,
			&privateTask.UserId,
			&privateTask.BoardId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		privateTasks = append(privateTasks, privateTask)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(privateTasks) == 0 {
		return nil, repository.ErrPrivateTaskNotFound
	}

	return privateTasks, nil
}

func (repo *PrivateTaskRepository) GetPrivateTasksWithBoardId(ctx context.Context, boardId int) ([]model.PrivateTask, error) {
	const op = "internal.repository.postgres.privateTask.GetPrivateTasks"
	const getPrivateTasksStmt = `SELECT id, category_id, name, description, status, date_create, deadline, user_id, board_id
					FROM private_tasks WHERE board_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getPrivateTasksStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var privateTasks []model.PrivateTask
	for rows.Next() {
		var privateTask model.PrivateTask
		err = rows.Scan(
			&privateTask.Id,
			&privateTask.CategoryId,
			&privateTask.Name,
			&privateTask.Description,
			&privateTask.Status,
			&privateTask.DateCreate,
			&privateTask.Deadline,
			&privateTask.UserId,
			&privateTask.BoardId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		privateTasks = append(privateTasks, privateTask)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(privateTasks) == 0 {
		return nil, repository.ErrPrivateTaskNotFound
	}

	return privateTasks, nil
}

func (repo *PrivateTaskRepository) GetPrivateTasksWithCategoryId(ctx context.Context, categoryId int) ([]model.PrivateTask, error) {

	const op = "internal.repository.postgres.privateTask.GetPrivateTasks"
	const getPrivateTasksStmt = `SELECT id, category_id, name, description, status, date_create, deadline, user_id, board_id
					FROM private_tasks WHERE category_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getPrivateTasksStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, categoryId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var privateTasks []model.PrivateTask
	for rows.Next() {
		var privateTask model.PrivateTask
		err = rows.Scan(
			&privateTask.Id,
			&privateTask.CategoryId,
			&privateTask.Name,
			&privateTask.Description,
			&privateTask.Status,
			&privateTask.DateCreate,
			&privateTask.Deadline,
			&privateTask.UserId,
			&privateTask.BoardId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		privateTasks = append(privateTasks, privateTask)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(privateTasks) == 0 {
		return nil, repository.ErrPrivateTaskNotFound
	}

	return privateTasks, nil
}

func (repo *PrivateTaskRepository) GetPrivateTaskWithId(ctx context.Context, id int) (*model.PrivateTask, error) {
	const op = "internal.repository.postgres.privateTask.GetPrivateTask"
	const getPrivateTaskStmt = `SELECT id, category_id, name, description, status, date_create, deadline, user_id, board_id
					FROM private_tasks WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getPrivateTaskStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var privateTask model.PrivateTask
	err = stmt.QueryRowContext(ctx, id).Scan(
		&privateTask.Id,
		&privateTask.CategoryId,
		&privateTask.Name,
		&privateTask.Description,
		&privateTask.Status,
		&privateTask.DateCreate,
		&privateTask.Deadline,
		&privateTask.UserId,
		&privateTask.BoardId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrPrivateTaskNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &privateTask, nil
}

func (repo *PrivateTaskRepository) UpdatePrivateTask(ctx context.Context, privateTask *model.PrivateTask) error {
	const op = "internal.repository.postgres.privateTask.UpdatePrivateTask"
	const updatePrivateTaskStmt = `UPDATE private_tasks SET category_id = $1, name = $2, description = $3, status = $4, date_create = $5, deadline = $6, board_id = $7 WHERE id = $8`

	stmt, err := repo.db.PrepareContext(ctx, updatePrivateTaskStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, privateTask.CategoryId, privateTask.Name, privateTask.Description, privateTask.Status, privateTask.DateCreate, privateTask.Deadline, privateTask.BoardId, privateTask.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrPrivateTaskNotFound
	}

	return nil
}

func (repo *PrivateTaskRepository) DeletePrivateTask(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.privateTask.DeletePrivateTask"
	const deletePrivateTaskStmt = `DELETE FROM private_tasks WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deletePrivateTaskStmt)
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
		return repository.ErrPrivateTaskNotFound
	}

	return nil
}
