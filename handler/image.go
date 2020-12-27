package handler

import (
	"github.com/Leonardo-Antonio/api-convert-to-b64/helper"
	"github.com/Leonardo-Antonio/api-convert-to-b64/model"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type Image struct{}

func NewImage() *Image {
	return &Image{}
}

func (i *Image) EncryptB64(c echo.Context) error {
	imageData := model.Image{}
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
	defer file.Close()
	b64, err := ioutil.ReadAll(file)
	if err != nil {
		response := helper.ResponseJSON(helper.MESSAGE, err.Error(), false, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	imageData.Src = b64
	imageData.FileName = fileHeader.Filename
	imageData.Size = fileHeader.Size
	imageData.Header = fileHeader.Header

	response := helper.ResponseJSON(
		helper.MESSAGE,
		"the image has been successfully converted",
		false,
		imageData,
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
