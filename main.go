package main

import (
	"context"
	"crypto/tls"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
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
	ctx := context.Background()
	db, err := fgconn.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
