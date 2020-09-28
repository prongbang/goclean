package handler_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/prongbang/goclean/api/v1/promotion/di"
	"github.com/prongbang/goclean/api/v1/promotion/gateway/handler"
	"github.com/prongbang/goclean/api/v1/promotion/model"
	"github.com/stretchr/testify/assert"
)

var handle handler.PromotionHandler
var e *echo.Echo

func init() {
	handle = handler.NewPromotionHandler(di.ProvidePromotionUseCase())
	e = echo.New()
}

const (
	promotionJson = `{"id":1,"code":"sd-promo","name":"Sunday promotion","priority":4,"exclusive":false,"used":0,"couponBased":false,"rules":[],"actions":[],"createdAt":"2017-02-28T12:05:12+0100","updatedAt":"2017-02-28T12:05:13+0100","channels":[],"_links":{"self":{"href":"/api/v1/promotions/sd-promo"}}}`
)

func setupHttpRequest(method, target string, body string) *http.Request {
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

func TestAddSuccess(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", promotionJson)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handle.Add(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestAddBadRequest(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", "{}")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := handle.Add(ctx)
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestAddBodyInvalidBadRequest(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", "{id:'1}")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := handle.Add(ctx)
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestAddNoIdBadRequest(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := handle.Add(ctx)
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestAddAndGetAllSuccess(t *testing.T) {
	TestAddSuccess(t)

	req := setupHttpRequest(echo.GET, "/api/v1/promotion", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handle.GetAll(ctx)) {
		var promotion []model.Promotion
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body)
		assert.NoError(t, json.Unmarshal([]byte(rec.Body.String()), &promotion))
		assert.Equal(t, 1, len(promotion))
		assert.Equal(t, 1, promotion[0].ID)
	}
}

func TestAddAndGetByIdSuccess(t *testing.T) {
	TestAddSuccess(t)

	req := setupHttpRequest(echo.GET, "/api/v1/", promotionJson)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/promotion/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, handle.Get(ctx)) {
		var promotion model.Promotion
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body)
		assert.NoError(t, json.Unmarshal([]byte(rec.Body.String()), &promotion))
		assert.Equal(t, 1, promotion.ID)
	}
}

func TestAddAndGetByIdBadRequest(t *testing.T) {
	TestAddSuccess(t)

	req := setupHttpRequest(echo.GET, "/api/v1/", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/promotion/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("")

	if assert.NoError(t, handle.Get(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body)
	}
}

func TestAddAndGetByIdNotFound(t *testing.T) {
	TestAddBadRequest(t)

	req := setupHttpRequest(echo.GET, "/api/v1/", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/promotion/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")

	if assert.NoError(t, handle.Get(ctx)) {
		log.Println(rec.Body)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotEmpty(t, rec.Body)
	}
}
