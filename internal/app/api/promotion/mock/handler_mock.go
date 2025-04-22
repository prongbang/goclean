package mock

import (
	"github.com/labstack/echo/v4"
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/stretchr/testify/mock"
)

type HandlerMock struct {
	mock.Mock
}

func (h *HandlerMock) Create(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}

func (h *HandlerMock) Find(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}

func (h *HandlerMock) FindList(c echo.Context) error {
	args := h.Called(c)
	return args.Error(0)
}

func NewHandler() promotion.Handler {
	return &HandlerMock{}
}
