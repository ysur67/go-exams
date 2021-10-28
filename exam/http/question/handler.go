package questhttp

import (
	"net/http"

	"example.com/exams/exam"
	"example.com/exams/models"
	"github.com/gin-gonic/gin"
)

type Question struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	ExamId string `json:"examId"`
	Number int    `json:"nubmer"`
	Body   string `json:"body"`
}

type Handler struct {
	questionUseCase exam.QuestionUseCase
}

func NewHandler(usecase exam.QuestionUseCase) *Handler {
	return &Handler{
		questionUseCase: usecase,
	}
}

func (handler *Handler) Get(ctx *gin.Context) {
	questions, err := handler.questionUseCase.GetQuestions(ctx, ctx.Param("examId"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	jsonQuestions := toQuestions(questions)
	ctx.JSON(
		http.StatusCreated,
		jsonQuestions,
	)
}

func (handler *Handler) Create(ctx *gin.Context) {
	var question Question
	err := ctx.ShouldBind(&question)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	exam, err := handler.questionUseCase.GetExam(ctx, question.ExamId)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	err = handler.questionUseCase.CreateQuestion(ctx, toModel(question, exam))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

}

func toQuestions(questions []models.Question) []Question {
	out := make([]Question, len(questions))
	for index, quest := range questions {
		out[index] = toQuestion(quest)
	}
	return out
}

func toQuestion(quest models.Question) Question {
	return Question{
		Id:     quest.Id,
		ExamId: quest.Exam.Id,
		Title:  quest.Title,
		Body:   quest.Body,
		Number: quest.Number,
	}
}

func toModel(quest Question, exam models.Exam) models.Question {
	return models.Question{
		Id:     quest.Id,
		Title:  quest.Title,
		Exam:   exam,
		Body:   quest.Body,
		Number: quest.Number,
	}
}
