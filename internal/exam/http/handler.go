package http

import (
	"fmt"
	"net/http"
	"time"

	exam "example.com/internal"
	"example.com/models"
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
	examUseCase exam.ExamUseCase
}

func NewHandler(usecase exam.ExamUseCase) *Handler {
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
	jsonExams := toExams(exams)
	ctx.JSON(
		http.StatusOK,
		jsonExams,
	)
}

func (handler *Handler) GetDetail(ctx *gin.Context) {
	exam, err := handler.examUseCase.GetDetailExam(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		toExam(exam),
	)
}

func toExams(exams []models.Exam) []Exam {
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
