package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	exam "example.com/internal"
	answerhttp "example.com/internal/answer/http"
	answerRepo "example.com/internal/answer/repository/postgres"
	answerUseCase "example.com/internal/answer/usecase"
	examhttp "example.com/internal/exam/http"
	examRepo "example.com/internal/exam/repository/postgres"
	examUseCase "example.com/internal/exam/usecase"
	questionhttp "example.com/internal/question/http"
	questionRepo "example.com/internal/question/repository/postgres"
	questUseCase "example.com/internal/question/usecase"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type App struct {
	server *http.Server

	examUseCase     exam.ExamUseCase
	questionUseCase exam.QuestionUseCase
	answerUseCase   exam.AnswerUseCase
}

func NewApp() *App {
	db := initDb()
	ctx := context.Background()
	questionRepo := questionRepo.NewQuestionRepository(db)
	err := questionRepo.InitTables(ctx)
	if err != nil {
		panic(err)
	}
	examRepo := examRepo.NewExamRepository(db)
	answerRepo := answerRepo.NewAnswerRepository(db)
	answerRepo.InitTables(ctx)

	return &App{
		examUseCase:     examUseCase.NewExamUseCase(examRepo, questionRepo, answerRepo),
		questionUseCase: questUseCase.NewQuestoinUseCase(questionRepo, examRepo),
		answerUseCase:   answerUseCase.NewAnswerRepository(answerRepo, questionRepo),
	}
}

func (app *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")

	examhttp.RegisterEndPoints(api, app.examUseCase)
	questionhttp.RegisterEndPoints(api, app.questionUseCase)
	answerhttp.RegisterEndPoints(api, app.answerUseCase)

	// HTTP Server
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
