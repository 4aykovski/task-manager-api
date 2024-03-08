package model

import "time"

type PrivateTask struct {
	Id          int
	UserId      string
	DeskId      int
	Name        string
	Description string
	Status      bool
	DateCreate  time.Time
	Deadline    time.Time
}
