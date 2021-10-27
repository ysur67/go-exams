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
