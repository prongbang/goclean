package promotion

import (
	"github.com/labstack/echo/v4"
)

// Router is the interface
type Router interface {
	Initialize(e *echo.Echo)
}

type router struct {
	Handle Handler
}

// NewRouter is the function new promotion router
func NewRouter(handler Handler) Router {
	return &router{
		Handle: handler,
	}
}

func (r *router) Initialize(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/promotion", r.Handle.FindList)
	v1.POST("/promotion", r.Handle.Create)
	v1.GET("/promotion/:id", r.Handle.Find)
}
