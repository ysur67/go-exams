package models

type Question struct {
	Id     string
	Exam   Exam
	Number int
	Title  string
	Body   string
}

type QuestionDetail struct {
	Id      string
	Exam    Exam
	Number  int
	Title   string
	Body    string
	Answers []Answer
}
