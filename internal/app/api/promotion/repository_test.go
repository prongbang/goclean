package promotion_test

import (
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/pkg/database"
	"log"
	"testing"
)

var repo promotion.Repository

func init() {
	ds = promotion.NewDataSource(database.NewDrivers())
}

func TestRepository_Create(t *testing.T) {
	data := mockPromotion()

	if err := repo.Create(&data); err != nil {
		t.Error("Cannot add promotion")
	}

	if data.ID == 1 {
		log.Println("Create promotion success")
	}
}

func TestRepository_CreateAndGetSlice(t *testing.T) {
	TestRepository_Create(t)

	data, err := repo.FindList()
	if err != nil {
		t.Error("Cannot get all promotion")
	}

	if len(data) == 0 {
		t.Error("Data not found")
	}
}

func TestRepository_CreateAndGet(t *testing.T) {
	TestRepository_Create(t)

	data, err := repo.Find(1)
	if err != nil {
		t.Error("Cannot get promotion by id")
	}

	if data.ID == 0 {
		t.Error("Data not found")
	}

}
