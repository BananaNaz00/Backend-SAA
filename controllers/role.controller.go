package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

func DataRole(c *echo.Context) error {
	result, err := models.DataRole()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func SimpanRole(c *echo.Context) error {
	var input models.InputRole

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format data JSON salah: " + err.Error(),
		})
	}

	result, err := models.SimpanRole(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateRole(c *echo.Context) error {
	id := c.FormValue("id")
	namaRole := c.FormValue("namaRole")
	deskripsiRole := c.FormValue("deskripsiRole")

	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateRole(conv_id, namaRole, deskripsiRole)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteRole(c *echo.Context) error {
	id := c.Param("id")

	conv_id, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteRole(conv_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
