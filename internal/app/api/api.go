package api

import (
	"github.com/labstack/echo/v4"
)

type API interface {
	Register()
}

type apis struct {
	Routers
}

func (a *apis) Register() {
	e := echo.New()

	// Routes
	a.Initialize(e)

	// Listener
	e.Logger.Fatal(e.Start(":1323"))
}

func New(routers Routers) API {
	return &apis{
		Routers: routers,
	}
}
