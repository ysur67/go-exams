package examhttp

import (
	"example.com/exams/exam"
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.RouterGroup, usecase exam.ExamUseCase) {
	h := NewHandler(usecase)
	exams := router.Group("/exams")
	{
		exams.GET("", h.Get)
		exams.GET("/:id", h.GetDetail)
	}
}
