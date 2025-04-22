package promotion_test

import (
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/pkg/database"
	"log"
	"testing"
)

var usecase promotion.UseCase

func init() {
	ds = promotion.NewDataSource(database.NewDrivers())
}

func TestUseCase_CreateThenError(t *testing.T) {
	data := mockPromotion()
	if err := usecase.Create(&data); err != nil {
		t.Error("Cannot add promotion")
	}

	if data.ID == 1 {
		log.Println("Create promotion success")
	}
}

func TestUseCase_Create(t *testing.T) {
	data := promotion.Promotion{}
	if err := usecase.Create(&data); err == nil {
		t.Error("Cannot add promotion")
	}

	if data.ID == 0 {
		log.Println("Create promotion success")
	}
}

func TestUseCase_CreateAndGetSlice(t *testing.T) {
	TestUseCase_Create(t)

	data, err := usecase.FindList()
	if err != nil {
		t.Error("Cannot get all promotion")
	}

	if len(data) == 0 {
		t.Error("Data not found")
	}
}

func TestUseCase_CreateAndGet(t *testing.T) {
	TestUseCase_Create(t)

	data, err := usecase.Find(1)
	if err != nil {
		t.Error("Cannot get promotion by id")
	}

	if data.ID == 0 {
		t.Error("Data not found")
	}
}
