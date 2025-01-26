package main

import (
	auth "Tour/internal/auth_microservice/app"
	gate "Tour/internal/gateway/app"

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
	auth := auth.New()
	gate := gate.New()

	go auth.Run()
	gate.Run()
}
