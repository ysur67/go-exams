package http

import (
	exam "example.com/internal"
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.RouterGroup, usecase exam.QuestionUseCase) {
	h := NewHandler(usecase)
	quests := router.Group("/exams/:examId/questions")
	{
		quests.GET("", h.Get)
		quests.POST("", h.Create)
	}
}
