package model

import "time"

type User struct {
	Id           string
	Login        string
	Email        string
	Password     string
	IsAdmin      bool
	About        string
	Language     string
	Theme        string
	RegisteredAt time.Time
	LastSeenAt   time.Time
}
