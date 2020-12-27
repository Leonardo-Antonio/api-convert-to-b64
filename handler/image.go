package handler

import (
	"fmt"
	"github.com/Leonardo-Antonio/api-convert-to-b64/helper"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type Image struct{}

func NewImage() *Image {
	return &Image{}
}

func (i *Image) EncryptB64(c echo.Context) error {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		response := helper.ResponseJSON(helper.MESSAGE, err.Error(), false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	file, err := fileHeader.Open()
	if err != nil {
		response := helper.ResponseJSON(helper.MESSAGE, err.Error(), false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	b64, err := ioutil.ReadAll(file)
	if err != nil {
		response := helper.ResponseJSON(helper.MESSAGE, err.Error(), false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	fmt.Println(fileHeader.Filename)
	response := helper.ResponseJSON(
		helper.MESSAGE,
		"the image has been successfully converted",
		false,
		b64,
	)
	return c.JSON(http.StatusOK, response)
}

func (i *Image) Test(c echo.Context) error {
	response := helper.ResponseJSON(
		helper.MESSAGE,
		"the image has been successfully converted",
		false,
		"Test - heroku deployd",
	)
	return c.JSON(http.StatusOK, response)
}
