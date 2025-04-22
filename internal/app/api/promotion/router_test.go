package promotion_test

import (
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/app/api/promotion/mock"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

var router promotion.Router
var e *echo.Echo

func init() {
	e = echo.New()
	router = promotion.NewRouter(mock.NewHandler())
	router.Initialize(e)
}

func TestRoute(t *testing.T) {
	r := e.Router()

	// Routes
	path := "/api/v1/promotion"
	r.Add(http.MethodGet, path, func(echo.Context) error {
		return nil
	})
	r.Add(http.MethodGet, path+"/:id", func(echo.Context) error {
		return nil
	})

	c := e.NewContext(nil, nil)

	r.Find(http.MethodGet, path, c)
	assert.Equal(t, "", c.Param(""))

	r.Find(http.MethodGet, path+"/1", c)
	assert.Equal(t, "1", c.Param("id"))
}
