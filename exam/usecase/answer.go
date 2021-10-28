package usecase

import (
	"context"

	"example.com/exams/exam"
	"example.com/exams/models"
)

type AnswerUseCase struct {
	answerRepo exam.AnswerRepository
}

func NewAnswerRepository(repo exam.AnswerRepository) *AnswerUseCase {
	return &AnswerUseCase{
		answerRepo: repo,
	}
}

func (useCase *AnswerUseCase) GetAnswers(ctx context.Context, questId string) ([]models.Answer, error) {
	return useCase.answerRepo.GetAnswers(ctx, questId)
}
