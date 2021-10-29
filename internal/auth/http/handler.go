package http

import (
	"net/http"

	exam "example.com/internal"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase exam.UserUseCase
}

func NewHandler(useCase exam.UserUseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type registerInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (handler *Handler) Register(ctx *gin.Context) {
	inp := new(registerInput)
	if err := ctx.BindJSON(inp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := handler.useCase.Register(ctx.Request.Context(), *toModel(*inp)); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err,
		)
		return
	}
	ctx.Status(http.StatusCreated)
}

func toModel(register registerInput) *models.LoginParam {
	return &models.LoginParam{
		Username: register.Username,
		Password: register.Password,
	}
}
