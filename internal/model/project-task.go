package model

import "time"

type ProjectTask struct {
	Id          int
	Name        string
	Description string
	Status      bool
	DateCreate  time.Time
	Deadline    time.Time
	BoardId     int
	CategoryId  int
	ProjectId   int
}
