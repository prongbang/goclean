package swagger

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Router is the interface
type Router interface {
	Initialize(e *echo.Echo)
}

type router struct {
}

// NewRouter is the function new promotion router
func NewRouter() Router {
	return &router{}
}

func (r *router) Initialize(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
