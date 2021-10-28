package usecase

import (
	"context"

	"example.com/exams/exam"
	"example.com/exams/models"
)

type QuestionUseCase struct {
	questRepo exam.QuestionRepository
	examRepo  exam.ExamRepository
}

func NewQuestoinUseCase(questRepo exam.QuestionRepository, examRepo exam.ExamRepository) *QuestionUseCase {
	return &QuestionUseCase{
		questRepo: questRepo,
		examRepo:  examRepo,
	}
}

func (useCase *QuestionUseCase) GetQuestions(ctx context.Context, examId string) ([]models.Question, error) {
	return useCase.questRepo.GetQuestions(ctx, examId)
}

func (useCase *QuestionUseCase) CreateQuestion(ctx context.Context, question models.Question) error {
	return useCase.questRepo.CreateQuestion(ctx, question)
}

func (useCase *QuestionUseCase) GetExam(ctx context.Context, examId string) (models.Exam, error) {
	return useCase.examRepo.GetDetailExam(ctx, examId)
}
