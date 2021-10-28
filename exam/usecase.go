package exam

import (
	"context"

	"example.com/exams/models"
)

type ExamUseCase interface {
	GetExams(ctx context.Context) ([]models.Exam, error)
	GetDetailExam(ctx context.Context, examId string) (models.Exam, error)
}

type QuestionUseCase interface {
	GetQuestions(ctx context.Context, examId string) ([]models.Question, error)
	GetExam(ctx context.Context, examId string) (models.Exam, error)
	CreateQuestion(ctx context.Context, question models.Question) error
}

type AnswerUseCase interface {
	GetAnswers(ctx context.Context, questId string) ([]models.Answer, error)
}
