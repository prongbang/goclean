package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prongbang/goclean/api/core"
	"github.com/prongbang/goclean/api/v1/promotion/domain"
	"github.com/prongbang/goclean/api/v1/promotion/model"
)

// PromotionHandler is the interface
type PromotionHandler interface {
	Add(c echo.Context) error
	GetAll(c echo.Context) error
	Get(c echo.Context) error
}

// Encapsulated Implementation of Handler Interface
type promotionHandler struct {
	UseCase domain.PromotionUseCase
}

// NewPromotionHandler is the function new promotion handler
func NewPromotionHandler(useCase domain.PromotionUseCase) PromotionHandler {
	return &promotionHandler{
		UseCase: useCase,
	}
}

// @Tags promotions
// @Summary Create promotion
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param promotion body model.Promotion true "Promotion"
// @Success 201 {object} model.Promotion
// @Failure 400 {object} core.Error
// @Security APIKeyAuth
// @Router /api/v1/promotion [post]
func (h *promotionHandler) Add(c echo.Context) error {
	var promotion model.Promotion
	if err := c.Bind(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, core.Error{
			Message: "Data not found",
		})
	}

	if err := h.UseCase.Add(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, core.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, promotion)
}

// @Tags promotions
// @Summary List promotion
// @Accept json
// @Produce json
// @Success 200 {object} model.Promotion
// @Security APIKeyAuth
// @Router /api/v1/promotion [get]
func (h *promotionHandler) GetAll(c echo.Context) error {

	promotions, _ := h.UseCase.GetAll()

	return c.JSON(http.StatusOK, promotions)
}

// @Tags promotions
// @Summary Get a promotion
// @Accept json
// @Produce json
// @Param id path int true "Promotion ID"
// @Success 200 {object} model.Promotion
// @Failure 400 {object} core.Error
// @Failure 404 {object} core.Error
// @Security APIKeyAuth
// @Router /api/v1/promotion/{id} [get]
func (h *promotionHandler) Get(c echo.Context) error {
	if id := c.Param("id"); id != "" {
		idInt, iErr := strconv.Atoi(id)
		if idInt > 0 && iErr == nil {
			promotion, err := h.UseCase.Get(idInt)
			if err == nil {
				return c.JSON(http.StatusOK, promotion)
			}
			return c.JSON(http.StatusNotFound, core.Error{
				Message: err.Error(),
			})
		}
	}
	return c.JSON(http.StatusBadRequest, core.Error{
		Message: "Required parameter id",
	})
}
