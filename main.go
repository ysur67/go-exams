package main

import (
	"fmt"
	"os"

	"example.com/exams/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("asd")
	app := server.NewApp()
	app.Run(os.Getenv("app-port"))
}
