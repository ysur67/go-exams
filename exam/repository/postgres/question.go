package postgres

import (
	"context"
	"strconv"

	"example.com/exams/models"
	"github.com/uptrace/bun"
)

const (
	QUESTION = "question"
)

type Question struct {
	Id     int64
	ExamId int64
	Exam   Exam `bun:"rel:has-one"`
	Number int
	Title  string
	Body   string
}

type QuestionRepository struct {
	db *bun.DB
}

func NewQuestionRepository(db *bun.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (repo *QuestionRepository) GetQuestions(ctx context.Context, examId string) ([]models.Question, error) {
	questions := make([]Question, 0)
	err := repo.db.NewSelect().
		Table(QUESTION).
		Scan(ctx, &questions)
	return toQuestions(questions), err
}

func toQuestions(questoins []Question) []models.Question {
	out := make([]models.Question, len(questoins))
	for index, quest := range questoins {
		out[index] = toQuestion(quest)
	}
	return out
}

func toQuestion(quest Question) models.Question {
	return models.Question{
		Id:     strconv.Itoa(int(quest.Id)),
		Number: quest.Number,
		Title:  quest.Title,
		Body:   quest.Body,
		Exam:   toExam(quest.Exam),
	}
}
