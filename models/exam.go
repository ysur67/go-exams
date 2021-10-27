package models

import "time"

type Exam struct {
	Id         string
	Title      string
	StartDate  time.Time
	FinishDate time.Time
	IsActive   bool
}
