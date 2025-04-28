package api

import (
	"github.com/labstack/echo/v4"
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/app/api/swagger"
)

type Routers interface {
	Initialize(c *echo.Echo)
}

type routers struct {
	PromotionRoute promotion.Router
	SwaggerRoute   swagger.Router
}

func (r *routers) Initialize(c *echo.Echo) {
	r.PromotionRoute.Initialize(c)
	r.SwaggerRoute.Initialize(c)
}

func NewRouters(
	promotionRoute promotion.Router,
	swaggerRoute swagger.Router,
) Routers {
	return &routers{
		PromotionRoute: promotionRoute,
		SwaggerRoute:   swaggerRoute,
	}
}
