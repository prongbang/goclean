package di

import (
	"github.com/prongbang/goclean/api/v1/promotion/data/datasource"
	"github.com/prongbang/goclean/api/v1/promotion/data/repository"
	"github.com/prongbang/goclean/api/v1/promotion/domain"
	"github.com/prongbang/goclean/api/v1/promotion/gateway/handler"
)

// Provide Database
func ProvideDatabaseHelper() datasource.DatabaseHelper {

	return datasource.GetDatabaseMock()
}

// Provide DataSource
func ProvidePromotionDataSource() datasource.PromotionDataSource {

	return datasource.NewPromotionDataSource(ProvideDatabaseHelper())
}

// Provide Repository
func ProvidePromotionRepository() repository.PromotionRepository {

	return repository.NewPromotionRepository(ProvidePromotionDataSource())
}

// Provide UseCase
func ProvidePromotionUseCase() domain.PromotionUseCase {

	return domain.NewPromotionUseCase(ProvidePromotionRepository())
}

// Provide Handler
func ProvidePromotionHandler() handler.PromotionHandler {

	return handler.NewPromotionHandler(ProvidePromotionUseCase())
}
