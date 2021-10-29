package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	exam "example.com/internal"
	answerhttp "example.com/internal/answer/http"
	answerRepo "example.com/internal/answer/repository/postgres"
	answerUseCase "example.com/internal/answer/usecase"
	authhttp "example.com/internal/auth/http"
	userRepo "example.com/internal/auth/repository/postgres"
	userUseCase "example.com/internal/auth/usecase"
	examhttp "example.com/internal/exam/http"
	examRepo "example.com/internal/exam/repository/postgres"
	examUseCase "example.com/internal/exam/usecase"
	questionhttp "example.com/internal/question/http"
	questionRepo "example.com/internal/question/repository/postgres"
	questUseCase "example.com/internal/question/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/xhit/go-str2duration/v2"
)

type App struct {
	server *http.Server

	examUseCase     exam.ExamUseCase
	questionUseCase exam.QuestionUseCase
	answerUseCase   exam.AnswerUseCase
	userUseCase     exam.UserUseCase
}

func NewApp() *App {
	db := initDb()
	ctx := context.Background()
	questionRepo := questionRepo.NewQuestionRepository(db)
	if err := questionRepo.InitTables(ctx); err != nil {
		panic(err)
	}
	examRepo := examRepo.NewExamRepository(db)
	if err := examRepo.InitTables(ctx); err != nil {
		panic(err)
	}
	answerRepo := answerRepo.NewAnswerRepository(db)
	if err := answerRepo.InitTables(ctx); err != nil {
		panic(err)
	}
	userRepo := userRepo.NewUserRepository(db)
	if err := userRepo.InitTables(ctx); err != nil {
		panic(err)
	}
	shouldLoadFixtures := getOptions()
	if shouldLoadFixtures {
		if err := loadFixtures(db); err != nil {
			panic(err)
		}
	}
	ttlDuration, err := str2duration.ParseDuration(os.Getenv("token-ttl"))
	if err != nil {
		panic(err)
	}
	return &App{
		examUseCase:     examUseCase.NewExamUseCase(examRepo, questionRepo, answerRepo),
		questionUseCase: questUseCase.NewQuestoinUseCase(questionRepo, examRepo),
		answerUseCase:   answerUseCase.NewAnswerUseCase(answerRepo, questionRepo),
		userUseCase: userUseCase.NewUserUseCase(
			userRepo,
			os.Getenv("hash-salt"),
			[]byte(os.Getenv("signin-key")),
			ttlDuration,
		),
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	authhttp.RegisterHttpEndpoints(router, app.userUseCase)

	authMiddleware := authhttp.NewAuthMiddleware(app.userUseCase)
	api := router.Group("/api", authMiddleware.Handle)
	examhttp.RegisterEndPoints(api, app.examUseCase)
	questionhttp.RegisterEndPoints(api, app.questionUseCase)
	answerhttp.RegisterEndPoints(api, app.answerUseCase)

	app.server = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.server.Shutdown(ctx)
}

func initDb() *bun.DB {
	dsn := "postgres://" +
		os.Getenv("db-username") +
		":" + os.Getenv("db-password") +
		"@" + os.Getenv("db-address") +
		"/" + os.Getenv("db-name") + "?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func getOptions() bool {
	cmd := os.Args[1:]
	if len(cmd) < 1 {
		return false
	}
	shouldLoadFixtures := cmd[0]
	return strings.Contains(shouldLoadFixtures, "load-fixtures")
}

func loadFixtures(db *bun.DB) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB),
		testfixtures.Dialect("postgres"),
		testfixtures.Paths(
			"configs/fixtures/users.yml",
			"configs/fixtures/exams.yml",
			"configs/fixtures/questions.yml",
			"configs/fixtures/answers.yml",
		),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		panic(err)
	}
	return fixtures.Load()
}
