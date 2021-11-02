package http

import (
	exam "example.com/internal"
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.RouterGroup, usecase exam.ExamUseCase) {
	h := NewHandler(usecase)
	exams := router.Group("/exams")
	{
		exams.GET("", h.Get)
		exams.GET("/:examId", h.GetDetail)
		exams.DELETE("/:examId", h.Delete)
		exams.POST("", h.Create)
	}
}
