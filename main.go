package main

import (
	"github.com/Leonardo-Antonio/api-convert-to-b64/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	//"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	router.Image(e)
	port := os.Getenv("POST")
	if port == "" {
		port = ":8080"
	}
	e.Logger.Fatal(e.Start(port))
}
