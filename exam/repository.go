package exam

import (
	"context"

	"example.com/exams/models"
)

type Repository interface {
	InitTables(ctx context.Context)
	GetExams(ctx context.Context) ([]models.Exam, error)
}
