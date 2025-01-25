package services

import (
	"Tour/internal/requests"
	"errors"

	"github.com/sirupsen/logrus"
)

type (
	Repository interface {
		SelectUser(requests.RegistrationRequest) error
		InsertUser(requests.RegistrationRequest) error
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

	if err := s.repository.InsertUser(user); err != nil {
		logrus.Error("Ошибка при добавлении пользователя: ", err)
		return errors.New("не удалось зарегистрировать пользователя")
	}

	logrus.Info("Регистрация успешна - ", user)
	return nil

}

func (s *Services) Login(user requests.LoginRequest) error {
	if user.Name == "1" {
		return errors.New("d")
	}
	// "дописать"
	return nil
}
