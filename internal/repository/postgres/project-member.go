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

type ProjectMemberRepository struct {
	db *postgres.Postgres
}

func NewProjectMemberRepository(db *postgres.Postgres) *ProjectMemberRepository {
	return &ProjectMemberRepository{db: db}
}

func (repo *ProjectMemberRepository) InsertProjectMember(ctx context.Context, projectMember model.ProjectMember) error {
	const op = "internal.repository.postgres.ProjectMemberRepository.InsertProjectMember"
	const insertProjectMemberStmt = `INSERT INTO project_members (project_id, user_id, status) VALUES ($1, $2, $3)`

	stmt, err := repo.db.PrepareContext(ctx, insertProjectMemberStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, projectMember.ProjectId, projectMember.UserId, projectMember.Status)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "unique_violation" {
				return repository.ErrProjectMemberAlreadyExists
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *ProjectMemberRepository) GetProjectMembers(ctx context.Context, projectId int) ([]model.ProjectMember, error) {
	const op = "internal.repository.postgres.ProjectMemberRepository.GetProjectMembers"
	const getProjectMemberStmt = `SELECT project_id, user_id, status FROM project_members WHERE project_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, getProjectMemberStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var projectMembers []model.ProjectMember
	for rows.Next() {
		var projectMember model.ProjectMember
		err := rows.Scan(
			&projectMember.ProjectId,
			&projectMember.UserId,
			&projectMember.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		projectMembers = append(projectMembers, projectMember)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(projectMembers) == 0 {
		return nil, repository.ErrProjectMemberNotFound
	}

	return projectMembers, nil
}

func (repo *ProjectMemberRepository) GetProjectMember(ctx context.Context, projectId int, userId string) (*model.ProjectMember, error) {
	const op = "internal.repository.postgres.ProjectMemberRepository.GetProjectMember"
	const getProjectMemberStmt = `SELECT project_id, user_id, status FROM project_members WHERE project_id = $1 AND user_id = $2`

	stmt, err := repo.db.PrepareContext(ctx, getProjectMemberStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var projectMember model.ProjectMember
	err = stmt.QueryRowContext(ctx, projectId, userId).Scan(
		&projectMember.ProjectId,
		&projectMember.UserId,
		&projectMember.Status,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrProjectMemberNotFound
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &projectMember, nil
}

func (repo *ProjectMemberRepository) GetUserProjects(ctx context.Context, userId string) ([]model.ProjectMember, error) {
	const op = "internal.repository.postgres.ProjectMemberRepository.GetUserProjects"
	const sql = `SELECT project_id, user_id, status FROM project_members WHERE user_id = $1`

	stmt, err := repo.db.PrepareContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var userProjects []model.ProjectMember
	for rows.Next() {
		var userProject model.ProjectMember
		err := rows.Scan(
			userProject.ProjectId,
			userProject.UserId,
			userProject.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		userProjects = append(userProjects, userProject)
	}

	if len(userProjects) == 0 {
		return nil, repository.ErrProjectMemberNotFound
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return userProjects, nil
}

func (repo *ProjectMemberRepository) DeleteProjectMember(ctx context.Context, projectId int, userId string) error {
	const op = "internal.repository.postgres.ProjectMemberRepository.DeleteProjectMember"
	const deleteProjectMemberStmt = `DELETE FROM project_members WHERE project_id = $1 AND user_id = $2`

	stmt, err := repo.db.PrepareContext(ctx, deleteProjectMemberStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, projectId, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrProjectMemberNotFound
	}

	return nil
}

func (repo *ProjectMemberRepository) UpdateProjectMember(ctx context.Context, projectMember model.ProjectMember) error {
	const op = "internal.repository.postgres.ProjectMemberRepository.UpdateProjectMember"
	const updateProjectMemberStmt = `UPDATE project_members SET status = $1 WHERE project_id = $2 AND user_id = $3`

	stmt, err := repo.db.PrepareContext(ctx, updateProjectMemberStmt)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, projectMember.Status, projectMember.ProjectId, projectMember.UserId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return repository.ErrProjectMemberNotFound
	}

	return nil
}
