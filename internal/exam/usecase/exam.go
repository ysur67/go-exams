package usecase

import (
	"context"

	exam "example.com/internal"
	"example.com/models"
)

type ExamUseCase struct {
	examRepo exam.ExamRepository
}

func NewExamUseCase(repo exam.ExamRepository) *ExamUseCase {
	return &ExamUseCase{
		examRepo: repo,
	}
}

func (useCase *ExamUseCase) GetExams(ctx context.Context) ([]models.Exam, error) {
	return useCase.examRepo.GetExams(ctx)
}

func (useCase *ExamUseCase) GetDetailExam(ctx context.Context, examId string) (models.Exam, error) {
	return useCase.examRepo.GetDetailExam(ctx, examId)
}
