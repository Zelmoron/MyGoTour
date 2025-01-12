package endpoints

import (
	"Tour/internal/requests"
	"net/http"

	"github.com/go-playground/validator/v10"
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

func (e *Endpoints) Registration(c *fiber.Ctx) error {
	user := requests.RegistrationRequest{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Request error",
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Validation error",
		})
	}
	return c.Status(http.StatusOK).JSON("")

}
