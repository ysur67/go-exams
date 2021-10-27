package exam

import (
	"context"

	"example.com/exams/models"
)

type UseCase interface {
	GetExams(ctx context.Context) ([]models.Exam, error)
}
