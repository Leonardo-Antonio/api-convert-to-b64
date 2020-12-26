package router

import (
	"github.com/Leonardo-Antonio/image-to-b64/handler"
	"github.com/labstack/echo/v4"
)

func Image(e *echo.Echo) {
	hand := handler.NewImage()
	group := e.Group("/api/v1/image")
	group.POST("/encrypt", hand.EncryptB64)
}
