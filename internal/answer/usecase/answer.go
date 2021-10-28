package usecase

import (
	"context"

	exam "example.com/internal"
	"example.com/models"
)

type AnswerUseCase struct {
	answerRepo   exam.AnswerRepository
	questionRepo exam.QuestionRepository
}

func NewAnswerUseCase(answerRepo exam.AnswerRepository, questionRepo exam.QuestionRepository) *AnswerUseCase {
	return &AnswerUseCase{
		answerRepo:   answerRepo,
		questionRepo: questionRepo,
	}
}

func (useCase *AnswerUseCase) GetAnswers(ctx context.Context, questId string) ([]models.Answer, error) {
	return useCase.answerRepo.GetAnswers(ctx, questId)
}

func (useCase *AnswerUseCase) CreateAnswer(ctx context.Context, answer models.Answer) error {
	return useCase.answerRepo.CreateAnswer(ctx, answer)
}

func (useCase *AnswerUseCase) GetQuestion(ctx context.Context, questId string) (models.Question, error) {
	return useCase.questionRepo.GetQuestion(ctx, questId)
}
