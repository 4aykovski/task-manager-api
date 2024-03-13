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
)
