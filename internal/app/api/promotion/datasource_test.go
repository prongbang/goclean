package promotion_test

import (
	"encoding/json"
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/pkg/database"
	"log"
	"testing"
)

var ds promotion.DataSource

func init() {
	ds = promotion.NewDataSource(database.NewDrivers())
}

func mockPromotion() promotion.Promotion {
	jsonString := `
	{
		"id": 1,
		"code": "sd-promo",
		"name": "Sunday promotion",
		"priority": 4,
		"exclusive": false,
		"used": 0,
		"couponBased": false,
		"rules": [],
		"actions": [],
		"createdAt": "2017-02-28T12:05:12+0100",
		"updatedAt": "2017-02-28T12:05:13+0100",
		"channels": [],
		"_links": {
			"self": {
				"href": "\/api\/v1\/promotions\/sd-promo"
			}
		}
	}`
	var data promotion.Promotion
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		log.Println("Cannot Unmarshal")
	}
	return data
}

func TestDataSource_Create(t *testing.T) {
	data := mockPromotion()
	if err := ds.Create(&data); err != nil {
		t.Error("Cannot add promotion")
	}

	if data.ID == 1 {
		log.Println("Create promotion success")
	}
}

func TestDataSource_GetThenError(t *testing.T) {
	data := promotion.Promotion{}
	if err := ds.Create(&data); err == nil {
		t.Error("Cannot add promotion")
	}

	if data.ID == 0 {
		log.Println("Create promotion success")
	}
}

func TestDataSource_CreateAndGetSlice(t *testing.T) {
	TestDataSource_Create(t)

	data, err := ds.FindList()
	if err != nil {
		t.Error("Cannot get all promotion")
	}

	if len(data) == 0 {
		t.Error("Data not found")
	}
}

func TestDataSource_CreateAndGet(t *testing.T) {
	TestDataSource_Create(t)

	data, err := ds.Find(1)
	if err != nil {
		t.Error("Cannot get promotion by id")
	}

	if data.ID == 0 {
		t.Error("Data not found")
	}

}

func TestAddAndGetByIdNotfound(t *testing.T) {
	TestDataSource_Create(t)

	data, err := ds.Find(2)
	if err == nil {
		t.Error("Found data get promotion by id 2")
	}

	if data.ID != 0 {
		t.Error("Data found id 2")
	}

}

func BenchmarkAdd(b *testing.B) {
	data := mockPromotion()
	// run function b.N times
	for n := 0; n < b.N; n++ {
		_ = ds.Create(&data)
	}
}
