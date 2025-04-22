package promotion

import (
	"fmt"
	"github.com/prongbang/goclean/internal/pkg/database"
)

// DataSource is the interface
type DataSource interface {
	Create(data *Promotion) error
	Find(id int) (Promotion, error)
	FindList() ([]Promotion, error)
}

// Encapsulated Implementation of DataSource Interface
type dataSource struct {
	Store database.Store
}

func (repo *dataSource) Create(data *Promotion) error {
	if data.ID > 0 {
		return repo.Store.Set(fmt.Sprint(data.ID), *data)
	}
	return fmt.Errorf("cannot add data: %s", "is empty")
}

func (repo *dataSource) FindList() ([]Promotion, error) {
	data := make([]Promotion, 0)
	values, err := repo.Store.Values()
	if err != nil {
		return data, err
	}
	for _, value := range values {
		data = append(data, value.(Promotion))
	}
	return data, nil
}

func (repo *dataSource) Find(id int) (Promotion, error) {
	data, err := repo.Store.Get(fmt.Sprint(id))
	if err != nil {
		return Promotion{}, err
	}
	return data.(Promotion), err
}

// NewDataSource is the new promotion datasource
func NewDataSource(drivers database.Drivers) DataSource {
	return &dataSource{
		Store: drivers.InMemoryStore,
	}
}
