package main

import (
	"github.com/Leonardo-Antonio/api-convert-to-b64/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	router.Image(e)
	port := os.Getenv("POST")
	e.Logger.Fatal(e.Start(port))
}
