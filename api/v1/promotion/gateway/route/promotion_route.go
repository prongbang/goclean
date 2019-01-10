package route

import (
	"github.com/labstack/echo"
	"github.com/prongbang/goclean/api/v1/promotion/gateway/handler"
)

type PromotionRoute interface {
	Initial(e *echo.Echo)
}

type promotionRoute struct {
	Handle handler.PromotionHandler
}

func NewPromotionRoute(handler handler.PromotionHandler) PromotionRoute {
	return &promotionRoute{
		Handle: handler,
	}
}

func (r *promotionRoute) Initial(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/promotion", r.Handle.GetAll)
	v1.POST("/promotion", r.Handle.Add)
	v1.GET("/promotion/:id", r.Handle.Get)
}
