package exam

import (
	"context"

	"example.com/exams/models"
)

type ExamRepository interface {
	InitTables(ctx context.Context)
	GetExams(ctx context.Context) ([]models.Exam, error)
	GetDetailExam(ctx context.Context, examId string) (models.Exam, error)
}

type QuestionRepository interface {
	InitTables(ctx context.Context) error
	GetQuestions(ctx context.Context, examId string) ([]models.Question, error)
	CreateQuestion(ctx context.Context, question models.Question) error
}

type AnswerRepository interface {
	InitTables(ctx context.Context) error
	GetAnswers(ctx context.Context, questId string) ([]models.Answer, error)
}
