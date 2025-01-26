package gateway

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}), logger.New(), recover.New())
	app.Post("/auth/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:8081"+c.OriginalURL())
	})

	// app.Get("/user/*", func(c *fiber.Ctx) error {
	// 	return proxy.Do(c, "http://user-service:8082"+c.OriginalURL())
	// })

	log.Println("API Gateway is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
