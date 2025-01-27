package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type App struct {
	app *fiber.App
}

func New() *App {
	a := &App{}
	a.app = fiber.New()
	a.router()
	return a
}

func (a *App) router() {
	a.app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// a.app.Use(logger.New())
	// a.app.Use(recover.New())

	// Прокси маршруты
	a.app.Post("/auth/*", func(c *fiber.Ctx) error {
		target := "http://localhost:8081" + c.OriginalURL()
		err := proxy.Do(c, target)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "Service unavailable",
			})
		}
		return nil
	})
	// app.Get("/user/*", func(c *fiber.Ctx) error {
	// 	return proxy.Do(c, "http://user-service:8082"+c.OriginalURL())
	// })

}

func (a *App) Run() {
	log.Println("API Gateway is running on port 8080")
	log.Fatal(a.app.Listen(":8080"))
}
