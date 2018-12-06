package controllers

import (
	"github.com/kurohi/shareskill/models"
	"net/http"
	"github.com/labstack/echo"
)


func GetUsers(c echo.Context) error {
	result := models.GetUsers()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
