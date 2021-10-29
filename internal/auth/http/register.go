package http

import (
	auth "example.com/internal"
	"github.com/gin-gonic/gin"
)

func RegisterHttpEndpoints(router *gin.Engine, usecase auth.UserUseCase) {
	handler := NewHandler(usecase)
	authEndPoints := router.Group("/auth")
	{
		authEndPoints.POST("/register", handler.Register)
	}
}
