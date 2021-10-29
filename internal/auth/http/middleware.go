package http

import (
	"net/http"
	"strings"

	exam "example.com/internal"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	useCase exam.UserUseCase
}

func NewAuthMiddleware(useCase exam.UserUseCase) *AuthMiddleware {
	return &AuthMiddleware{
		useCase: useCase,
	}
}

func (middleware *AuthMiddleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := middleware.useCase.ParseToken(ctx.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusUnauthorized

		ctx.AbortWithStatus(status)
		return
	}

	ctx.Set(exam.CtxUserKey, user)
}
