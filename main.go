package main

import (
	"fmt"
	"github.com/Leonardo-Antonio/api-convert-to-b64/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	//"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	fmt.Println("PORTTTTTT no POST", port)
	e := echo.New()
	e.Use(middleware.CORS())
	router.Image(e)
	e.Logger.Fatal(e.Start(port))
}
