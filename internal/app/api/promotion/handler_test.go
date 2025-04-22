package promotion_test

import (
	"encoding/json"
	"github.com/prongbang/goclean/internal/app/api/promotion"
	"github.com/prongbang/goclean/internal/app/api/promotion/mock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var handle promotion.Handler
var mockUseCase *mock.UseCaseMock

func init() {
	mockUseCase = &mock.UseCaseMock{}
	handle = promotion.NewHandler(mockUseCase)
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

func TestHandler_Create(t *testing.T) {
	promo := promotion.Promotion{ID: 1, Name: "Sunday promotion"}
	mockUseCase.On("Create", &promo).Return(nil)

	req := setupHttpRequest(echo.POST, "/api/v1/promotion", promotionJson)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handle.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}

	mockUseCase.AssertExpectations(t)
}

func TestHandler_CreateAndGet(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", "{}")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := handle.Create(ctx)
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestHandler_CreateAndBodyInvalidBadRequest(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", "{id:'1}")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := handle.Create(ctx)
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestHandler_CreateAndNoIdBadRequest(t *testing.T) {
	req := setupHttpRequest(echo.POST, "/api/v1/promotion", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	res := handle.Create(ctx)
	if assert.NoError(t, res) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestHandler_CreateAndGetAllSuccess(t *testing.T) {
	TestHandler_Create(t)

	req := setupHttpRequest(echo.GET, "/api/v1/promotion", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handle.FindList(ctx)) {
		var promotions []promotion.Promotion
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body)
		assert.NoError(t, json.Unmarshal([]byte(rec.Body.String()), &promotions))
		assert.Equal(t, 1, len(promotions))
		assert.Equal(t, 1, promotions[0].ID)
	}
}

func TestHandler_CreateAndGetByIdSuccess(t *testing.T) {
	TestHandler_Create(t)

	req := setupHttpRequest(echo.GET, "/api/v1/", promotionJson)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/promotion/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, handle.Find(ctx)) {
		var p promotion.Promotion
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body)
		assert.NoError(t, json.Unmarshal([]byte(rec.Body.String()), &p))
		assert.Equal(t, 1, p.ID)
	}
}

func TestHandler_CreateAndGetByIdBadRequest(t *testing.T) {
	TestHandler_Create(t)

	req := setupHttpRequest(echo.GET, "/api/v1/", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/promotion/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("")

	if assert.NoError(t, handle.Find(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotEmpty(t, rec.Body)
	}
}

func TestHandler_CreateAndGetByIdNotFound(t *testing.T) {
	TestHandler_Create(t)

	req := setupHttpRequest(echo.GET, "/api/v1/", "")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/promotion/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")

	if assert.NoError(t, handle.Find(ctx)) {
		log.Println(rec.Body)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotEmpty(t, rec.Body)
	}
}
