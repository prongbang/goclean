package promotion

import (
	"github.com/prongbang/goclean/pkg/core"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Handler is the interface
type Handler interface {
	Create(c echo.Context) error
	Find(c echo.Context) error
	FindList(c echo.Context) error
}

// Encapsulated Implementation of Handler Interface
type handler struct {
	UseCase UseCase
}

// NewHandler is the function new promotion handler
func NewHandler(useCase UseCase) Handler {
	return &handler{
		UseCase: useCase,
	}
}

// Create
// @Tags promotions
// @Summary Create promotion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param promotion body Promotion true "Promotion"
// @Success 201 {object} Promotion
// @Failure 400 {object} core.Error
// @Security APIKeyAuth
// @Router /api/v1/promotion [post]
func (h *handler) Create(c echo.Context) error {
	var promotion Promotion
	if err := c.Bind(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, core.Error{Message: "Data not found"})
	}
	if err := h.UseCase.Create(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, core.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, promotion)
}

// FindList
// @Tags promotions
// @Summary List promotion
// @Accept json
// @Produce json
// @Success 200 {object} Promotion
// @Security APIKeyAuth
// @Router /api/v1/promotion [get]
func (h *handler) FindList(c echo.Context) error {
	promotions, _ := h.UseCase.FindList()
	return c.JSON(http.StatusOK, promotions)
}

// Find
// @Tags promotions
// @Summary Find a promotion
// @Accept json
// @Produce json
// @Param id path int true "Promotion ID"
// @Success 200 {object} Promotion
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security APIKeyAuth
// @Router /api/v1/promotion/{id} [get]
func (h *handler) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	promotion, err := h.UseCase.Find(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, core.Error{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, promotion)
}
