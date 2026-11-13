package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func DataTingkatUrgensi(c *echo.Context) error {
	result, err := models.DataTingkatUrgensi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func SimpanTingkatUrgensi(c *echo.Context) error {
	namaTingkatUrgensi := c.FormValue("namaTingkatUrgensi")

	result, err := models.SimpanTingkatUrgensi(namaTingkatUrgensi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
