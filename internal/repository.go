package exam

import (
	"context"

	"example.com/models"
)

type ExamRepository interface {
	InitTables(ctx context.Context) error
	GetExams(ctx context.Context) ([]models.Exam, error)
	GetDetailExam(ctx context.Context, examId string) (models.Exam, error)
	CreateExam(ctx context.Context, exam models.Exam) error
}

type QuestionRepository interface {
	InitTables(ctx context.Context) error
	GetQuestion(ctx context.Context, id string) (models.Question, error)
	GetQuestions(ctx context.Context, examId string) ([]models.Question, error)
	CreateQuestion(ctx context.Context, question models.Question) error
}

type AnswerRepository interface {
	InitTables(ctx context.Context) error
	GetAnswers(ctx context.Context, questId string) ([]models.Answer, error)
	CreateAnswer(ctx context.Context, answer models.Answer) error
}

type UserRepository interface {
	InitTables(ctx context.Context) error
	CreateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, authParams models.LoginParam) (models.User, error)
}
