package exam

import (
	"context"

	"example.com/models"
)

type ExamUseCase interface {
	GetExams(ctx context.Context) ([]models.Exam, error)
	GetDetailExam(ctx context.Context, examId string) (models.ExamDetail, error)
	CreateExam(ctx context.Context, exam models.Exam) error
	RemoveExam(ctx context.Context, params models.RemoveExamParams) error
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

const CtxUserKey = "user"

type UserUseCase interface {
	Register(ctx context.Context, authParams models.LoginParam) error
	Login(ctx context.Context, authParams models.LoginParam) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
