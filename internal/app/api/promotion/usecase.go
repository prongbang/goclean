package promotion

import (
	"errors"
)

// UseCase is the interface
type UseCase interface {
	Create(data *Promotion) error
	Find(id int) (Promotion, error)
	FindList() ([]Promotion, error)
}

// Encapsulated Implementation of UseCase Interface
type useCase struct {
	Repo Repository
}

// NewUseCase is the function new promotion use case
func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}

func (uc *useCase) Create(data *Promotion) error {
	if err := uc.Repo.Create(data); err == nil {
		promo, err := uc.Find(data.ID)
		*data = promo

		return err
	}

	return errors.New("not found")
}

func (uc *useCase) FindList() ([]Promotion, error) {

	return uc.Repo.FindList()
}

func (uc *useCase) Find(id int) (Promotion, error) {

	return uc.Repo.Find(id)
}
