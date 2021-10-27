package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/exams/exam"
	examhttp "example.com/exams/exam/http/exam"
	questhttp "example.com/exams/exam/http/question"
	"example.com/exams/exam/repository/postgres"
	"example.com/exams/exam/usecase"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type App struct {
	server *http.Server

	examUseCase     exam.ExamUseCase
	questionUseCase exam.QuestionUseCase
}

func NewApp() *App {
	db := initDb()

	return &App{
		examUseCase:     usecase.NewExamUseCase(postgres.NewExamRepository(db)),
		questionUseCase: usecase.NewQuestoinUseCase(postgres.NewQuestionRepository(db)),
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
	questhttp.RegisterEndPoints(api, app.questionUseCase)

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
