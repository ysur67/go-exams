package main

import (
	"os"

	"example.com/exams/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := server.NewApp()
	app.Run(os.Getenv("app-port"))
}
