package main

import (
	"github.com/prongbang/goclean"
	_ "github.com/prongbang/goclean/docs"
	"github.com/prongbang/goclean/internal/pkg/database"
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
	drivers := database.NewDrivers()
	apis := goclean.NewAPI(drivers)
	apis.Register()
}
