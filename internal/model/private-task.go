package model

import "time"

type PrivateTask struct {
	Id          int
	CategoryId  int
	Name        string
	Description string
	Status      bool
	DateCreate  time.Time
	Deadline    time.Time
	UserId      string
	BoardId     int
}
