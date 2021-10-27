package models

type Question struct {
	Id     string
	Exam   Exam
	Number int
	Title  string
	Body   string
}
