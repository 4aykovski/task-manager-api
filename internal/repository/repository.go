package repository

import "errors"

var (
	ErrUserAlreadyExists            = errors.New("user already exists")
	ErrUserNotFound                 = errors.New("user not found")
	ErrProjectAlreadyExists         = errors.New("project already exists")
	ErrProjectNotFound              = errors.New("user not found")
	ErrProjectMemberAlreadyExists   = errors.New("project member already exists")
	ErrProjectMemberNotFound        = errors.New("project members not found")
	ErrRefreshSessionAlreadyExists  = errors.New("refresh session already exists")
	ErrRefreshSessionNotFound       = errors.New("refresh session not found")
	ErrPrivateBoardAlreadyExists    = errors.New("private board already exists")
	ErrPrivateBoardNotFound         = errors.New("private board not found")
	ErrPrivateCategoryAlreadyExists = errors.New("private category already exists")
	ErrPrivateCategoryNotFound      = errors.New("private category not found")
	ErrPrivateTaskAlreadyExists     = errors.New("private task already exists")
	ErrPrivateTaskNotFound          = errors.New("private task not found")
	ErrProjectBoardAlreadyExists    = errors.New("project board already exists")
	ErrProjectBoardNotFound         = errors.New("project board not found")
	ErrProjectCategoryAlreadyExists = errors.New("project category already exists")
	ErrProjectCategoryNotFound      = errors.New("project category not found")
	ErrProjectTaskAlreadyExists     = errors.New("project task already exists")
	ErrProjectTasksNotFound         = errors.New("project tasks not found")
	ErrProjectTaskNotFound          = errors.New("project task not found")
)

// TODO: add unique violation error on update to every repository
