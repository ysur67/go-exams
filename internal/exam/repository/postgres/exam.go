package postgres

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"example.com/models"
	"github.com/uptrace/bun"
)

const (
	EXAM = "exam"
)

type Exam struct {
	Id         int64
	Title      string
	Startdate  time.Time
	Finishdate time.Time
	Isactive   bool
}

type ExamRepository struct {
	db *bun.DB
}

func NewExamRepository(db *bun.DB) *ExamRepository {
	return &ExamRepository{
		db: db,
	}
}

func (repo *ExamRepository) InitTables(ctx context.Context) {
	repo.db.NewCreateTable().Model(&Exam{}).Table("exam").Temp().IfNotExists().Varchar(300).Exec(ctx)
}

func (repo *ExamRepository) GetExams(ctx context.Context) ([]models.Exam, error) {
	var exams = make([]Exam, 0)
	err := repo.db.NewSelect().
		Table(EXAM).
		Scan(ctx, &exams)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return toExams(exams), nil
}

func (repo *ExamRepository) GetDetailExam(ctx context.Context, examId string) (models.Exam, error) {
	exam := Exam{}
	err := repo.db.NewSelect().
		Table(EXAM).
		Where("id = ?", examId).
		Scan(ctx, &exam)
	return ToExam(exam), err
}

func toExams(exams []Exam) []models.Exam {
	out := make([]models.Exam, len(exams))
	for index, exam := range exams {
		out[index] = ToExam(exam)
	}
	return out
}

func ToExam(exam Exam) models.Exam {
	return models.Exam{
		Id:         strconv.Itoa(int(exam.Id)),
		Title:      exam.Title,
		StartDate:  exam.Startdate,
		FinishDate: exam.Finishdate,
		IsActive:   exam.Isactive,
	}
}
