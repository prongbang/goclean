package datasource

import (
	"fmt"

	"github.com/prongbang/goclean/api/v1/promotion/model"
)

// GetDatabaseMock is the Mock database in memory
func GetDatabaseMock() DatabaseHelper {
	return DatabaseHelper{
		Store: make(map[int]model.Promotion),
	}
}

// DatabaseHelper is the struct
type DatabaseHelper struct {
	Store map[int]model.Promotion
}

// PromotionDataSource is the interface
type PromotionDataSource interface {
	Add(data *model.Promotion) error
	GetAll() ([]model.Promotion, error)
	Get(id int) (model.Promotion, error)
}

// Encapsulated Implementation of DataSource Interface
type promotionDataSource struct {
	Db DatabaseHelper
}

// NewPromotionDataSource is the new promotion datasource
func NewPromotionDataSource(db DatabaseHelper) PromotionDataSource {
	return &promotionDataSource{
		Db: db,
	}
}

func (repo *promotionDataSource) Add(data *model.Promotion) error {

	if data.ID > 0 {

		repo.Db.Store[data.ID] = *data

		return nil
	}

	return fmt.Errorf("Cannot add data: %s", "is empty")
}

func (repo *promotionDataSource) GetAll() ([]model.Promotion, error) {

	data := []model.Promotion{}
	for _, value := range repo.Db.Store {
		data = append(data, value)
	}

	return data, nil
}

func (repo *promotionDataSource) Get(id int) (model.Promotion, error) {

	data := repo.Db.Store[id]
	if data.ID == id {

		return repo.Db.Store[id], nil
	}

	return model.Promotion{}, fmt.Errorf("Not found id: %d", id)
}
