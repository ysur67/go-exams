package usecase

import (
	"context"

	"example.com/exams/exam"
	"example.com/exams/models"
)

type QuestionUseCase struct {
	questRepo exam.QuestionRepository
}

func NewQuestoinUseCase(repo exam.QuestionRepository) *QuestionUseCase {
	return &QuestionUseCase{
		questRepo: repo,
	}
}

func (useCase *QuestionUseCase) GetQuestions(ctx context.Context, examId string) ([]models.Question, error) {
	return useCase.questRepo.GetQuestions(ctx, examId)
}
