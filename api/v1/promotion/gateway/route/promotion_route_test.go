package route_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	e := echo.New()
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
