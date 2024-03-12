package repository

import "errors"

var (
	ErrUserAlreadyExists          = errors.New("user already exists")
	ErrUserNotFound               = errors.New("user not found")
	ErrProjectAlreadyExists       = errors.New("project already exists")
	ErrProjectNotFound            = errors.New("user not found")
	ErrProjectMemberAlreadyExists = errors.New("project member already exists")
	ErrProjectMemberNotFound      = errors.New("project members not found")
)
