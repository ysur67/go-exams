package questhttp

import (
	"example.com/exams/exam"
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.RouterGroup, usecase exam.QuestionUseCase) {
	h := NewHandler(usecase)
	quests := router.Group("/exams/:id/questions")
	{
		quests.GET("", h.Get)
	}
}
