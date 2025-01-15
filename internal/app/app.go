package app

import (
	"Tour/internal/endpoints"
	"Tour/internal/repository"
	"Tour/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	endpoints  *endpoints.Endpoints
	services   *services.Services
	repository *repository.Repository
	app        *fiber.App
}

func New() *App {

	a := &App{}

	db := repository.Connect()
	a.repository = repository.New(db)
	a.services = services.New(a.repository)
	a.endpoints = endpoints.New(a.services)
	a.app = fiber.New()
	a.routers()

	return a
}

func (a *App) routers() {
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}), logger.New(), recover.New())

	public := a.app.Group("")
	public.Post("/registration", a.endpoints.Registration)

}

func (a *App) Run() {
	a.app.Listen(":8080")
}
