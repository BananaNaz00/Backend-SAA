package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func DataPenerimaanBarang(c *echo.Context) error {
	result, err := models.DataPenerimaanBarang()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func SimpanPenerimaanBarang(c *echo.Context) error {
	var req models.InputPenerimaanBarang

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format request tidak valid: " + err.Error(),
		})
	}

	result, err := models.SimpanPenerimaanBarang(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePenerimaanBarang(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error()})
	}

	var req models.InputPenerimaanBarang
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	result, err := models.UpdatePenerimaanBarang(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePenerimaanBarang(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Id Tidak Valid"})
	}

	result, err := models.DeletePenerimaanBarang(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
