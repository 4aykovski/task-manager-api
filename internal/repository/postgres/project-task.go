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

type ProjectTaskRepository struct {
	db *postgres.Postgres
}

func NewProjectTaskRepository(db *postgres.Postgres) *ProjectTaskRepository {
	return &ProjectTaskRepository{db}
}

func (repo *ProjectTaskRepository) CreateProjectTask(ctx context.Context, projectTask *model.ProjectTask) error {
	const op = "internal.repository.postgres.projectTask.CreateProjectTask"
	const createProjectTaskStmt = `INSERT INTO project_tasks (name, description, deadline, board_id, category_id, project_id)
									VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := repo.db.PrepareContext(ctx, createProjectTaskStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		projectTask.Name,
		projectTask.Description,
		projectTask.Deadline,
		projectTask.BoardId,
		projectTask.CategoryId,
		projectTask.ProjectId,
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrProjectTaskAlreadyExists
			}
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *ProjectTaskRepository) GetProjectTasksWithProjectId(ctx context.Context, projectId string) ([]model.ProjectTask, error) {
	const op = "internal.repository.postgres.projectTask.GetProjectTasks"
	const getProjectTasksStmt = `SELECT id, name, description, status, date_create, deadline, board_id, category_id, project_id
									FROM project_tasks WHERE project_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectTasksStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var projectTasks []model.ProjectTask
	for rows.Next() {
		var projectTask model.ProjectTask
		err = rows.Scan(
			&projectTask.Id,
			&projectTask.Name,
			&projectTask.Description,
			&projectTask.Status,
			&projectTask.DateCreate,
			&projectTask.Deadline,
			&projectTask.BoardId,
			&projectTask.CategoryId,
			&projectTask.ProjectId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projectTasks = append(projectTasks, projectTask)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(projectTasks) == 0 {
		return nil, repository.ErrProjectTasksNotFound
	}

	return projectTasks, nil
}

func (repo *ProjectTaskRepository) GetProjectTasksWithBoardId(ctx context.Context, boardId int) ([]model.ProjectTask, error) {
	const op = "internal.repository.postgres.projectTask.GetProjectTasks"
	const getProjectTasksStmt = `SELECT id, name, description, status, date_create, deadline, board_id, category_id, project_id
									FROM project_tasks WHERE board_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectTasksStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, boardId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var projectTasks []model.ProjectTask
	for rows.Next() {
		var projectTask model.ProjectTask
		err = rows.Scan(
			&projectTask.Id,
			&projectTask.Name,
			&projectTask.Description,
			&projectTask.Status,
			&projectTask.DateCreate,
			&projectTask.Deadline,
			&projectTask.BoardId,
			&projectTask.CategoryId,
			&projectTask.ProjectId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projectTasks = append(projectTasks, projectTask)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(projectTasks) == 0 {
		return nil, repository.ErrProjectTasksNotFound
	}

	return projectTasks, nil
}

func (repo *ProjectTaskRepository) GetProjectTasksWithCategoryId(ctx context.Context, categoryId int) ([]model.ProjectTask, error) {
	const op = "internal.repository.postgres.projectTask.GetProjectTasks"
	const getProjectTasksStmt = `SELECT id, name, description, status, date_create, deadline, board_id, category_id, project_id
									FROM project_tasks WHERE category_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectTasksStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, categoryId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var projectTasks []model.ProjectTask
	for rows.Next() {
		var projectTask model.ProjectTask
		err = rows.Scan(
			&projectTask.Id,
			&projectTask.Name,
			&projectTask.Description,
			&projectTask.Status,
			&projectTask.DateCreate,
			&projectTask.Deadline,
			&projectTask.BoardId,
			&projectTask.CategoryId,
			&projectTask.ProjectId,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projectTasks = append(projectTasks, projectTask)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(projectTasks) == 0 {
		return nil, repository.ErrProjectTasksNotFound
	}

	return projectTasks, nil
}

func (repo *ProjectTaskRepository) GetProjectTaskWithId(ctx context.Context, id int) (*model.ProjectTask, error) {
	const op = "internal.repository.postgres.projectTask.GetProjectTask"
	const getProjectTaskStmt = `SELECT id, name, description, status, date_create, deadline, board_id, category_id, project_id
									FROM project_tasks WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectTaskStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var projectTask model.ProjectTask
	err = stmt.QueryRowContext(ctx, id).Scan(
		&projectTask.Id,
		&projectTask.Name,
		&projectTask.Description,
		&projectTask.Status,
		&projectTask.DateCreate,
		&projectTask.Deadline,
		&projectTask.BoardId,
		&projectTask.CategoryId,
		&projectTask.ProjectId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrProjectTaskNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &projectTask, nil
}

func (repo *ProjectTaskRepository) UpdateProjectTask(ctx context.Context, projectTask *model.ProjectTask) error {
	const op = "internal.repository.postgres.projectTask.UpdateProjectTask"
	const updateProjectTaskStmt = `UPDATE project_tasks SET name = $1, description = $2, status = $3, date_create = $4, deadline = $5, board_id = $6, category_id = $7, project_id = $8 WHERE id = $9`

	stmt, err := repo.db.PrepareContext(ctx, updateProjectTaskStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(
		ctx,
		projectTask.Name,
		projectTask.Description,
		projectTask.Status,
		projectTask.DateCreate,
		projectTask.Deadline,
		projectTask.BoardId,
		projectTask.CategoryId,
		projectTask.ProjectId,
		projectTask.Id,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrProjectTaskNotFound
	}

	return nil
}

func (repo *ProjectTaskRepository) DeleteProjectTask(ctx context.Context, id int) error {
	const op = "internal.repository.postgres.projectTask.DeleteProjectTask"
	const deleteProjectTaskStmt = `DELETE FROM project_tasks WHERE id = $1`

	stmt, err := repo.db.PrepareContext(ctx, deleteProjectTaskStmt)
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
		return repository.ErrProjectTaskNotFound
	}

	return nil
}
