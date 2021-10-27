package models

type Answer struct {
	Id        string
	Question  Question
	Title     string
	IsCorrect bool
}
