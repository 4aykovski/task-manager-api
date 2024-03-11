package repository

import "errors"

var (
	ErrUserAlreadyExists           = errors.New("user already exists")
	ErrUserNotFound                = errors.New("user not found")
	ErrProjectAlreadyExists        = errors.New("project already exists")
	ErrProjectNotFound             = errors.New("user not found")
	ErrRefreshSessionAlreadyExists = errors.New("refresh session already exists")
	ErrRefreshSessionNotFound      = errors.New("refresh session not found")
)
