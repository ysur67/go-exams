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

func (repo *ExamRepository) InitTables(ctx context.Context) error {
	_, err := repo.db.NewCreateTable().
		Model(&Exam{}).
		IfNotExists().
		Varchar(300).
		Exec(ctx)
	return err
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
	return toModels(exams), nil
}

func (repo *ExamRepository) GetDetailExam(ctx context.Context, examId string) (models.Exam, error) {
	exam := Exam{}
	err := repo.db.NewSelect().
		Model(&exam).
		Where("id = ?", examId).
		Scan(ctx, &exam)
	return ToModel(exam), err
}

func (repo *ExamRepository) CreateExam(ctx context.Context, exam models.Exam) error {
	dbExam := toExam(exam)
	_, err := repo.db.NewInsert().
		Model(&dbExam).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx)
	return err
}

func toModels(exams []Exam) []models.Exam {
	out := make([]models.Exam, len(exams))
	for index, exam := range exams {
		out[index] = ToModel(exam)
	}
	return out
}

func ToModel(exam Exam) models.Exam {
	return models.Exam{
		Id:         strconv.Itoa(int(exam.Id)),
		Title:      exam.Title,
		StartDate:  exam.Startdate,
		FinishDate: exam.Finishdate,
		IsActive:   exam.Isactive,
	}
}

func toExam(model models.Exam) Exam {
	modelId, err := strconv.Atoi(model.Id)
	if err != nil {
		panic(err)
	}
	return Exam{
		Id:         int64(modelId),
		Title:      model.Title,
		Startdate:  model.StartDate,
		Finishdate: model.FinishDate,
		Isactive:   model.IsActive,
	}
}
