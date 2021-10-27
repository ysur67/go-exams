package main

import (
	"crypto/tls"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun/driver/pgdriver"
)

type App struct {
	server *http.Server
}

func NewApp() *App {
	db := initDb()
	db.Ping()
}

func initDb() *sql.DB {
	godotenv.Load()
	fgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(os.Getenv("db-address")),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(os.Getenv("db-username")),
		pgdriver.WithPassword(os.Getenv("db-password")),
		pgdriver.WithDatabase(os.Getenv("db-name")),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(fgconn.Config().Addr)))
	return sqldb
}
