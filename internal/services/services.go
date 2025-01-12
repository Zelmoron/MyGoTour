package services

import (
	"Tour/internal/requests"
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
)

type (
	Repository interface {
		SelectUser(context.Context, chan<- bool, chan bool, requests.RegistrationRequest)
		InsertUser(context.Context, <-chan bool, chan<- bool, chan bool, requests.RegistrationRequest)
	}
	Services struct {
		repository Repository
	}
)

func New(repository Repository) *Services {
	return &Services{
		repository: repository,
	}
}

func (s *Services) Compilator() {

}

func (s *Services) Registration(user requests.RegistrationRequest) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan bool)
	ch2 := make(chan bool)
	errChan := make(chan bool)

	go s.repository.SelectUser(ctx, ch1, errChan, user)
	go s.repository.InsertUser(ctx, ch1, ch2, errChan, user)

	log.Info("Ждем ch2")
	select {
	case err := <-ch2:
		if err {
			logrus.Info("Регистрация успешна")
		}
		return nil
	case <-errChan:
		logrus.Error("Ошибка Sql запроса")
		cancel()
		return nil
	}

}
