package route_test

import (
	"net/http"
	"testing"

	"github.com/prongbang/goclean/api/v1/promotion/di"
	"github.com/prongbang/goclean/api/v1/promotion/gateway/route"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

var router route.PromotionRoute
var e *echo.Echo

func init() {
	router = route.NewPromotionRoute(di.ProvidePromotionHandler())
	e = echo.New()
	router.Initial(e)
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
