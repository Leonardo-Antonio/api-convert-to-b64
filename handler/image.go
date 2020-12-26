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
		fmt.Println(err)
	}
	fmt.Println(fileHeader.Filename)
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error")
	}
	response := helper.ResponseJSON(
		helper.MESSAGE,
		"the image has been successfully converted",
		false,
		data,
	)
	return c.JSON(http.StatusOK, response)
}
