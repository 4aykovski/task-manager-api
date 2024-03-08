package model

import "time"

type Project struct {
	Id          int
	Name        string
	Owner       string
	Description string
	CreatedAt   time.Time
}
