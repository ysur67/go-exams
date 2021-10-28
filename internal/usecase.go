package exam

import (
	"context"

	"example.com/models"
)

type ExamUseCase interface {
	GetExams(ctx context.Context) ([]models.Exam, error)
	GetDetailExam(ctx context.Context, examId string) (models.ExamDetail, error)
	CreateExam(ctx context.Context, exam models.Exam) error
}

type QuestionUseCase interface {
	GetQuestions(ctx context.Context, examId string) ([]models.Question, error)
	GetExam(ctx context.Context, examId string) (models.Exam, error)
	CreateQuestion(ctx context.Context, question models.Question) error
}

type AnswerUseCase interface {
	GetAnswers(ctx context.Context, questId string) ([]models.Answer, error)
	GetQuestion(ctx context.Context, questId string) (models.Question, error)
	CreateAnswer(ctx context.Context, answer models.Answer) error
}
