package http

import (
	"fmt"
	"net/http"
	"time"

	exam "example.com/internal"
	questionHttp "example.com/internal/question/http"
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

type DetailExam struct {
	Id         string                        `json:"id"`
	Title      string                        `json:"title"`
	StartDate  time.Time                     `json:"startDate"`
	FinishDate time.Time                     `json:"finishDate"`
	IsActive   bool                          `json:"isActive"`
	Questions  []questionHttp.QuestionDetail `json:"questions"`
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
	exam, err := handler.examUseCase.GetDetailExam(ctx, ctx.Param("examId"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		toDetailExam(exam),
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

func toDetailExam(exam models.ExamDetail) DetailExam {
	detailQuestions := make([]questionHttp.QuestionDetail, len(exam.Questions))
	for index, quest := range exam.Questions {
		detailQuestions[index] = questionHttp.ToDetailQuestion(quest)
	}
	return DetailExam{
		Id:         exam.Id,
		Title:      exam.Title,
		StartDate:  exam.StartDate,
		FinishDate: exam.FinishDate,
		IsActive:   exam.IsActive,
		Questions:  detailQuestions,
	}
}
