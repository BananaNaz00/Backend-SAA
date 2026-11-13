package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v5"
)

func DataActivityLog(c *echo.Context) error {
	result, err := models.DataActivityLog()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func SimpanActivityLog(c *echo.Context) error {
	var req models.InputActivityLog

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	result, err := models.SimpanActivityLog(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
