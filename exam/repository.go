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
