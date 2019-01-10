package di

import (
	"github.com/prongbang/goclean/api/v1/promotion/data/datasource"
	"github.com/prongbang/goclean/api/v1/promotion/data/repository"
	"github.com/prongbang/goclean/api/v1/promotion/domain"
	"github.com/prongbang/goclean/api/v1/promotion/gateway/handler"
)

// ProvideDatabaseHelper is the Provide Database
func ProvideDatabaseHelper() datasource.DatabaseHelper {

	return datasource.GetDatabaseMock()
}

// ProvidePromotionDataSource is the Provide DataSource
func ProvidePromotionDataSource() datasource.PromotionDataSource {

	return datasource.NewPromotionDataSource(ProvideDatabaseHelper())
}

// ProvidePromotionRepository is the Provide Repository
func ProvidePromotionRepository() repository.PromotionRepository {

	return repository.NewPromotionRepository(ProvidePromotionDataSource())
}

// ProvidePromotionUseCase is the Provide UseCase
func ProvidePromotionUseCase() domain.PromotionUseCase {

	return domain.NewPromotionUseCase(ProvidePromotionRepository())
}

// ProvidePromotionHandler is the Provide Handler
func ProvidePromotionHandler() handler.PromotionHandler {

	return handler.NewPromotionHandler(ProvidePromotionUseCase())
}
