package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func Asatuan(c *echo.Context) error {
	result, err := models.Asatuan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
