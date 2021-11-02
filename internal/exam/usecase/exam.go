package usecase

import (
	"context"

	exam "example.com/internal"
	"example.com/models"
)

type ExamUseCase struct {
	examRepo     exam.ExamRepository
	questionRepo exam.QuestionRepository
	answerRepo   exam.AnswerRepository
}

func NewExamUseCase(examRepo exam.ExamRepository, questRepo exam.QuestionRepository, answerRepo exam.AnswerRepository) *ExamUseCase {
	return &ExamUseCase{
		examRepo:     examRepo,
		questionRepo: questRepo,
		answerRepo:   answerRepo,
	}
}

func (useCase *ExamUseCase) GetExams(ctx context.Context) ([]models.Exam, error) {
	return useCase.examRepo.GetExams(ctx)
}

func (useCase *ExamUseCase) CreateExam(ctx context.Context, exam models.Exam) error {
	return useCase.examRepo.CreateExam(ctx, exam)
}

func (useCase *ExamUseCase) GetDetailExam(ctx context.Context, examId string) (models.ExamDetail, error) {
	emptyExam := models.ExamDetail{}
	exam, err := useCase.examRepo.GetDetailExam(ctx, examId)
	if err != nil {
		return emptyExam, err
	}
	questions, err := useCase.questionRepo.GetQuestions(ctx, examId)
	if err != nil {
		return emptyExam, err
	}
	detailQuestions := make([]models.QuestionDetail, len(questions))
	for index, question := range questions {
		detailQuestions[index] = models.QuestionDetail{
			Id:     question.Id,
			Title:  question.Title,
			Body:   question.Body,
			Exam:   exam,
			Number: question.Number,
		}
		answers, err := useCase.answerRepo.GetAnswers(ctx, question.Id)
		if err != nil {
			return emptyExam, err
		}
		detailQuestions[index].Answers = answers
	}
	return models.ExamDetail{
		Id:         exam.Id,
		Title:      exam.Title,
		StartDate:  exam.StartDate,
		FinishDate: exam.FinishDate,
		IsActive:   exam.IsActive,
		Questions:  detailQuestions,
	}, nil
}

func (useCase *ExamUseCase) RemoveExam(ctx context.Context, params models.RemoveExamParams) error {
	return useCase.examRepo.Remove(ctx, params.Id)
}
