package main

import (
	"github.com/Leonardo-Antonio/api-convert-to-b64/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	router.Image(e)
	e.Logger.Fatal(e.Start(":8080"))
}
