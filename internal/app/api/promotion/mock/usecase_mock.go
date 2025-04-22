package mock

import (
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/stretchr/testify/mock"
)

type UseCaseMock struct {
	mock.Mock
}

func (m *UseCaseMock) Create(data *promotion.Promotion) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *UseCaseMock) Find(id int) (promotion.Promotion, error) {
	args := m.Called(id)
	return args.Get(0).(promotion.Promotion), args.Error(1)
}

func (m *UseCaseMock) FindList() ([]promotion.Promotion, error) {
	args := m.Called()
	return args.Get(0).([]promotion.Promotion), args.Error(1)
}

func NewUseCase() promotion.UseCase {
	return &UseCaseMock{}
}
