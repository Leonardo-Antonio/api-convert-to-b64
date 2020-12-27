package main

import (
	"github.com/Leonardo-Antonio/api-convert-to-b64/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	//"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("POST IN: %s", port)
	e := echo.New()
	e.Use(middleware.CORS())
	router.Image(e)
	e.Logger.Fatal(e.Start(":" + port))
}
