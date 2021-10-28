package postgres

import (
	"context"
	"strconv"

	"example.com/exams/models"
	"github.com/uptrace/bun"
)

type Answer struct {
	Id         int64
	QuestionId int64
	Question   Question `bun:"rel:has-one"`
	Title      string
	IsCorrect  bool
}

type AnswerRepository struct {
	db *bun.DB
}

const (
	ANSWER = "answers"
)

func NewAnswerRepository(db *bun.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}

func (repo *AnswerRepository) InitTables(ctx context.Context) error {
	_, err := repo.db.NewCreateTable().
		Model((*Answer)(nil)).
		IfNotExists().
		Varchar(300).
		ForeignKey(`("question_id") REFERENCES "questions" ("id") ON DELETE CASCADE`).
		Exec(ctx)
	return err
}

func (repo *AnswerRepository) GetAnswers(ctx context.Context, questId string) ([]models.Answer, error) {
	answers := make([]Answer, 0)
	qId, err := strconv.Atoi(questId)
	if err != nil {
		panic(err)
	}
	err = repo.db.NewSelect().
		Table(ANSWER).
		Where("question_id = ?", int64(qId)).
		Scan(ctx, &answers)
	return toAnswersModel(answers), err
}

func toAnswersModel(answers []Answer) []models.Answer {
	out := make([]models.Answer, len(answers))
	for index, answ := range answers {
		out[index] = toAnswModel(answ)
	}
	return out
}

func toAnswModel(answer Answer) models.Answer {
	return models.Answer{
		Id:        strconv.Itoa(int(answer.Id)),
		Title:     answer.Title,
		Question:  toModel(answer.Question),
		IsCorrect: answer.IsCorrect,
	}
}
