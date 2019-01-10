package datasource_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/prongbang/goclean/api/v1/promotion/data/datasource"
	"github.com/prongbang/goclean/api/v1/promotion/di"
	"github.com/prongbang/goclean/api/v1/promotion/model"
)

var ds datasource.PromotionDataSource

func init() {
	ds = datasource.NewPromotionDataSource(di.ProvideDatabaseHelper())
}

func mockPromotion() model.Promotion {
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
	var data model.Promotion
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		log.Println("Cannot Unmarshal")
	}
	return data
}

func TestAdd(t *testing.T) {
	data := mockPromotion()
	if err := ds.Add(&data); err != nil {
		t.Error("Cannot add promotion")
	}

	if data.ID == 1 {
		log.Println("Add promotion success")
	}
}

func TestAddAndGetAll(t *testing.T) {
	TestAdd(t)

	data, err := ds.GetAll()
	if err != nil {
		t.Error("Cannot get all promotion")
	}

	if len(data) == 0 {
		t.Error("Data not found")
	}
}

func TestAddAndGetById(t *testing.T) {
	TestAdd(t)

	data, err := ds.Get(1)
	if err != nil {
		t.Error("Cannot get promotion by id")
	}

	if data.ID == 0 {
		t.Error("Data not found")
	}

}

func BenchmarkAdd(b *testing.B) {
	data := mockPromotion()
	// run function b.N times
	for n := 0; n < b.N; n++ {
		ds.Add(&data)
	}
}
