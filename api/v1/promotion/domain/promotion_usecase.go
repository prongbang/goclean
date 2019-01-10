package domain

import (
	"fmt"

	"github.com/prongbang/goclean/api/v1/promotion/data/repository"
	"github.com/prongbang/goclean/api/v1/promotion/model"
)

// PromotionUseCase is the interface
type PromotionUseCase interface {
	Add(data *model.Promotion) error
	GetAll() ([]model.Promotion, error)
	Get(id int) (model.Promotion, error)
}

// Encapsulated Implementation of UseCase Interface
type promotionUseCase struct {
	Repo repository.PromotionRepository
}

// NewPromotionUseCase is the function new promotion use case
func NewPromotionUseCase(repo repository.PromotionRepository) PromotionUseCase {
	return &promotionUseCase{
		Repo: repo,
	}
}

func (uc *promotionUseCase) Add(data *model.Promotion) error {
	if err := uc.Repo.Add(data); err == nil {
		promo, err := uc.Get(data.ID)
		*data = promo

		return err
	}

	return fmt.Errorf("Not found")
}

func (uc *promotionUseCase) GetAll() ([]model.Promotion, error) {

	return uc.Repo.GetAll()
}

func (uc *promotionUseCase) Get(id int) (model.Promotion, error) {

	return uc.Repo.Get(id)
}
