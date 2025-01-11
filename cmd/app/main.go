package main

import (
	"Tour/internal/app"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
	}
}
func main() {
	app := app.New()
	app.Run()
}
