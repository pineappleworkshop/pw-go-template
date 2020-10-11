package main

import (
	"strconv"

	"{{<service_name>}}/config"
	"{{<service_name>}}/services"
	"{{<service_name>}}/stores"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.InitConf()
	stores.InitDbs()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	initRoutes(e)

	go e.Logger.Fatal(e.Start(":" + strconv.Itoa(services.PORT)))
}

func initRoutes(e *echo.Echo) {
	e.GET("/health", services.HealthHandler)
	e.POST("/product", services.ProductHandler)
}
