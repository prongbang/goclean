package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prongbang/goclean/api/v1/promotion/di"
	"github.com/prongbang/goclean/api/v1/promotion/gateway/route"
	_ "github.com/prongbang/goclean/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title GoClean API
// @version 1.0
// @description This is a swagger for GoClean API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url https://prongbang.github.io
// @license.name MIT License
// @license.url https://github.com/prongbang/goclean/blob/master/LICENSE
// @host localhost:1323
// @BasePath /
// @securityDefinitions.apikey APIKeyAuth
// @in header
// @name X-API-KEY
func main() {
	e := echo.New()

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	route.NewPromotionRoute(di.ProvidePromotionHandler()).Initial(e)

	// Listener
	e.Logger.Fatal(e.Start(":1323"))
}
