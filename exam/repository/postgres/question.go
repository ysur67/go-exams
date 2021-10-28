package postgres

import (
	"context"
	"strconv"

	"example.com/exams/models"
	"github.com/uptrace/bun"
)

const (
	QUESTION = "questions"
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

func (repo *QuestionRepository) InitTables(ctx context.Context) error {
	_, err := repo.db.NewCreateTable().
		Model((*Question)(nil)).
		Table("question").
		IfNotExists().
		Varchar(300).
		ForeignKey(`("exam_id") REFERENCES "exam" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	return err
}

func (repo *QuestionRepository) GetQuestions(ctx context.Context, examId string) ([]models.Question, error) {
	questions := make([]Question, 0)
	err := repo.db.NewSelect().
		Table(QUESTION).
		Scan(ctx, &questions)
	return toModels(questions), err
}

func (repo *QuestionRepository) CreateQuestion(ctx context.Context, question models.Question) error {
	dbQuest := toQuestion(question)
	_, err := repo.db.NewInsert().
		Model(&dbQuest).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx)
	return err
}

func toModels(questoins []Question) []models.Question {
	out := make([]models.Question, len(questoins))
	for index, quest := range questoins {
		out[index] = toModel(quest)
	}
	return out
}

func toModel(quest Question) models.Question {
	return models.Question{
		Id:     strconv.Itoa(int(quest.Id)),
		Number: quest.Number,
		Title:  quest.Title,
		Body:   quest.Body,
		Exam:   toExam(quest.Exam),
	}
}

func toQuestion(quest models.Question) Question {
	qId, err := strconv.Atoi(quest.Id)
	if err != nil {
		panic(err)
	}
	examId, err := strconv.Atoi(quest.Exam.Id)
	if err != nil {
		panic(err)
	}
	return Question{
		Id:     int64(qId),
		ExamId: int64(examId),
		Number: quest.Number,
		Title:  quest.Title,
		Body:   quest.Body,
	}
}
