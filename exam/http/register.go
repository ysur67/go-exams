package http

import (
	"example.com/exams/exam"
	"github.com/gin-gonic/gin"
)

func RegisterEndPoints(router *gin.RouterGroup, usecase exam.UseCase) {
	h := NewHandler(usecase)
	exams := router.Group("/exams")
	{
		exams.GET("", h.Get)
	}
}
