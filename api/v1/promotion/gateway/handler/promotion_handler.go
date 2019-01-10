package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
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

func (h *promotionHandler) Add(c echo.Context) error {
	var promotion model.Promotion
	if err := c.Bind(&promotion); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Data not found!",
		})
	}

	if err := h.UseCase.Add(&promotion); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}

	return c.JSON(http.StatusCreated, promotion)
}

func (h *promotionHandler) GetAll(c echo.Context) error {

	promotions, err := h.UseCase.GetAll()
	if err == nil {
		return c.JSON(http.StatusOK, promotions)
	}
	return c.JSON(http.StatusOK, []model.Promotion{})
}

func (h *promotionHandler) Get(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		idInt, iErr := strconv.Atoi(id)
		if idInt > 0 && iErr == nil {
			promotion, err := h.UseCase.Get(idInt)
			if err == nil {
				return c.JSON(http.StatusOK, promotion)
			}
			return c.JSON(http.StatusNotFound, err)
		}
	}
	return c.JSON(http.StatusBadGateway, echo.Map{
		"message": "Required parameter id",
	})
}
