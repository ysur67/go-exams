package answerhttp

import (
	"net/http"

	"example.com/exams/exam"
	"example.com/exams/models"
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

func (handler Handler) Get(ctx *gin.Context) {
	answers, err := handler.answerUseCase.GetAnswers(ctx, ctx.Param("questionid"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
	}
	jsonAnswers := toAnswers(answers)
	ctx.JSON(
		http.StatusOK,
		jsonAnswers,
	)
}

func toAnswers(answers []models.Answer) []Answer {
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
