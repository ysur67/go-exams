package postgres

import (
	"context"
	"strconv"

	repositoryModels "example.com/internal/question/repository/postgres"
	"example.com/models"
	"github.com/uptrace/bun"
)

type Answer struct {
	Id         int64
	QuestionId int64
	Question   repositoryModels.Question `bun:"rel:has-one"`
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
	return toModels(answers), err
}

func (repo *AnswerRepository) CreateAnswer(ctx context.Context, answer models.Answer) error {
	dbAnswer := toAnswer(answer)
	_, err := repo.db.NewInsert().
		Model(&dbAnswer).
		On("CONFLICT (id) DO UPDATE").
		Exec(ctx)
	return err
}

func toModels(answers []Answer) []models.Answer {
	out := make([]models.Answer, len(answers))
	for index, answ := range answers {
		out[index] = toModel(answ)
	}
	return out
}

func toModel(answer Answer) models.Answer {
	return models.Answer{
		Id:        strconv.Itoa(int(answer.Id)),
		Title:     answer.Title,
		Question:  repositoryModels.ToModel(answer.Question),
		IsCorrect: answer.IsCorrect,
	}
}

func toAnswer(model models.Answer) Answer {
	modelId, err := strconv.Atoi(model.Id)
	if err != nil {
		panic(err)
	}
	modelQuestionId, err := strconv.Atoi(model.Question.Id)
	if err != nil {
		panic(err)
	}
	return Answer{
		Id:         int64(modelId),
		Title:      model.Title,
		QuestionId: int64(modelQuestionId),
		// Question:   repositoryModels.ToQuestion(model.Question),
		IsCorrect: model.IsCorrect,
	}
}
