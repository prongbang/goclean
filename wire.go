//go:build wireinject
// +build wireinject

package goclean

import (
	"github.com/google/wire"
	"github.com/prongbang/goclean/internal/app/api"
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/app/api/swagger"
	"github.com/prongbang/goclean/internal/pkg/database"
)

func NewAPI(drivers database.Drivers) api.API {
	wire.Build(
		api.NewRouters,
		api.New,
		promotion.ProviderSet,
		swagger.ProviderSet,
	)
	return nil
}
