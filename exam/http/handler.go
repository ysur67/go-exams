package http

import (
	"fmt"
	"net/http"
	"time"

	"example.com/exams/exam"
	"example.com/exams/models"
	"github.com/gin-gonic/gin"
)

type Exam struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	StartDate  time.Time `json:"startDate"`
	FinishDate time.Time `json:"finishDate"`
	IsActive   bool      `json:"isActive"`
}

type Handler struct {
	examUseCase exam.UseCase
}

func NewHandler(usecase exam.UseCase) *Handler {
	return &Handler{
		examUseCase: usecase,
	}
}

func (handler *Handler) Get(ctx *gin.Context) {
	exams, err := handler.examUseCase.GetExams(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}
	jsonExams := toReponse(exams)
	ctx.JSON(
		http.StatusOK,
		jsonExams,
	)
}

func toReponse(exams []models.Exam) []Exam {
	out := make([]Exam, len(exams))
	for index, exam := range exams {
		out[index] = toExam(exam)
	}
	return out
}

func toExam(exam models.Exam) Exam {
	return Exam{
		Id:         exam.Id,
		Title:      exam.Title,
		StartDate:  exam.StartDate,
		FinishDate: exam.FinishDate,
		IsActive:   exam.IsActive,
	}
}
