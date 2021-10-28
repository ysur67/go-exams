package answerhttp

import (
	"example.com/exams/exam"
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.RouterGroup, useCase exam.AnswerUseCase) {
	handler := NewHandler(useCase)
	answers := router.Group("/exams/:examId/questions/:questionid/answers")
	{
		answers.GET("", handler.Get)
	}
}
