package usecase

import (
	"context"

	"example.com/exams/exam"
	"example.com/exams/models"
)

type ExamUseCase struct {
	examRepo exam.Repository
}

func NewExamUseCase(repo exam.Repository) *ExamUseCase {
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
