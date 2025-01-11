package endpoints

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type (
	Services interface {
		Compilator()
	}
	Endpoints struct {
		services Services
	}
)

func New(services Services) *Endpoints {
	return &Endpoints{
		services: services,
	}
}
func (e *Endpoints) TestHadler(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"test": "test",
	})
}
