package model

import "time"

type ProjectTask struct {
	Id          int
	ProjectId   int
	Name        string
	Description string
	Status      bool
	DateCreate  time.Time
	Deadline    time.Time
}
