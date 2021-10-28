package http

import (
	"net/http"

	exam "example.com/internal"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

type Answer struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	QuestionId string `json:"questionId"`
	IsCorrect  bool   `json:"isCorrect"`
}

type Handler struct {
	answerUseCase exam.AnswerUseCase
}

func NewHandler(useCase exam.AnswerUseCase) *Handler {
	return &Handler{
		answerUseCase: useCase,
	}
}

func (handler *Handler) Get(ctx *gin.Context) {
	answers, err := handler.answerUseCase.GetAnswers(ctx, ctx.Param("questionid"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	jsonAnswers := ToAnswers(answers)
	ctx.JSON(
		http.StatusOK,
		jsonAnswers,
	)
}

func (handler *Handler) Create(ctx *gin.Context) {
	var answer Answer
	err := ctx.ShouldBind(&answer)
	if err != nil {
		panic(err)
	}
	question, err := handler.answerUseCase.GetQuestion(ctx, answer.QuestionId)
	if err != nil {
		panic(err)
	}
	err = handler.answerUseCase.CreateAnswer(ctx, toModel(answer, question))
	if err != nil {
		panic(err)
	}
	ctx.JSON(
		http.StatusOK,
		"ok",
	)
}

func ToAnswers(answers []models.Answer) []Answer {
	out := make([]Answer, len(answers))
	for index, answ := range answers {
		out[index] = toAnswer(answ)
	}
	return out
}

func toAnswer(answer models.Answer) Answer {
	return Answer{
		Id:         answer.Id,
		Title:      answer.Title,
		QuestionId: answer.Question.Id,
		IsCorrect:  answer.IsCorrect,
	}
}

func toModel(answer Answer, question models.Question) models.Answer {
	return models.Answer{
		Id:        answer.Id,
		Title:     answer.Title,
		Question:  question,
		IsCorrect: answer.IsCorrect,
	}
}
