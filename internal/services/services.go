package services

type (
	Repository interface {
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
