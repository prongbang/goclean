package repository

import (
	"github.com/prongbang/goclean/api/v1/promotion/data/datasource"
	"github.com/prongbang/goclean/api/v1/promotion/model"
)

// PromotionRepository is the interface
type PromotionRepository interface {
	Add(data *model.Promotion) error
	GetAll() ([]model.Promotion, error)
	Get(id int) (model.Promotion, error)
}

// Encapsulated Implementation of Repository Interface
type promotionRepository struct {
	Ds datasource.PromotionDataSource
}

// NewPromotionRepository is the function
func NewPromotionRepository(ds datasource.PromotionDataSource) PromotionRepository {
	return &promotionRepository{
		Ds: ds,
	}
}

func (repo *promotionRepository) Add(data *model.Promotion) error {

	return repo.Ds.Add(data)
}

func (repo *promotionRepository) GetAll() ([]model.Promotion, error) {

	return repo.Ds.GetAll()
}

func (repo *promotionRepository) Get(id int) (model.Promotion, error) {

	return repo.Ds.Get(id)
}
