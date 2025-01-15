package endpoints

import (
	"Tour/internal/requests"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	Services interface {
		Compilator()
		Registration(requests.RegistrationRequest) error
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
		logrus.Error(fmt.Sprintf("Ошибка при получении данных %v", err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Request error",
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		logrus.Error(fmt.Sprintf("Ошибка при валидации данных %v", err))
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Validation error",
		})
	}

	logrus.Info(fmt.Sprintf("Данные получены %s", user))

	err := e.services.Registration(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "BadRequest - Registration error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "OK - Registration success",
	})

}
