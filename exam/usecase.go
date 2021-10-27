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
}
