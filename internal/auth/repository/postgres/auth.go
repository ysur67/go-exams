package postgres

import (
	"context"
	"fmt"
	"strconv"

	"example.com/models"
	"github.com/uptrace/bun"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) InitTables(ctx context.Context) error {
	_, err := repo.db.NewCreateTable().
		Model((*User)(nil)).
		IfNotExists().
		Varchar(150).
		Exec(ctx)
	return err
}

func (repo *UserRepository) CreateUser(ctx context.Context, user models.User) error {
	dbUser := toUser(user)
	_, err := repo.db.NewInsert().
		Model(dbUser).
		On("CONFLICT (id) DO NOTHING").
		Exec(ctx)
	return err
}

func (repo *UserRepository) GetUser(ctx context.Context, authParams models.LoginParam) (models.User, error) {
	user := new(User)
	err := repo.db.NewSelect().
		Model(user).
		Where("username = ?", authParams.Username).
		Where("password = ?", authParams.Password).
		Limit(1).
		Scan(ctx)
	return *toModel(*user), err
}

func toUser(model models.User) *User {
	uid, err := strconv.Atoi(model.Id)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	return &User{
		Id:       int64(uid),
		Username: model.Username,
		Password: model.Password,
	}
}

func toModel(user User) *models.User {
	return &models.User{
		Id:       strconv.Itoa(int(user.Id)),
		Username: user.Username,
		Password: user.Password,
	}
}
