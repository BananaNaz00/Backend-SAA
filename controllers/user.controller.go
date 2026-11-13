package controllers

//
//import (
//	"myapp/db"
//	"myapp/models"
//	"net/http"
//	"strconv"
//	"strings"
//
//	"github.com/labstack/echo/v5"
//	"golang.org/x/crypto/bcrypt"
//)
//
//const (
//	defaultIDRole       = 2
//	defaultIDStatusUser = 1
//	tempValue           = ""
//)
//
//func DataUser(c *echo.Context) error {
//	result, err := models.DataUser()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//	return c.JSON(http.StatusOK, result)
//}
//
//func Register(c *echo.Context) error {
//	var req models.RegisterRequest
//	if err := c.Bind(&req); err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "format request tidak valid")
//	}
//
//	//Username minimal 8 karakter
//	req.Username = strings.TrimSpace(req.Username)
//	if len(req.Username) < 8 {
//		return echo.NewHTTPError(http.StatusBadRequest, "username minimal 8 karakter")
//	}
//
//	req.Email = strings.TrimSpace(req.Email)
//	if req.Email == "" {
//		return echo.NewHTTPError(http.StatusBadRequest, "Email harus diisi")
//	}
//
//	if len(req.Password) < 8 || len(req.Password) > 72 {
//		return echo.NewHTTPError(http.StatusBadRequest, "password harus 8-72 karakter")
//	}
//
//	con := db.CreateCon()
//
//	//untuk cek username apakah sudah digunakan
//	var existingID int
//	err := con.QueryRow("SELECT idUser FROM user WHERE username = ?", req.Username).Scan(&existingID)
//	if err == nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "username sudah digunakan")
//	}
//
//	//untuk cek email apakah sudah digunakan
//	err = con.QueryRow("SELECT idUser FROM user WHERE email = ?", req.Email).Scan(&existingID)
//	if err == nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "email sudah digunakan")
//	}
//
//	//Untuk Hash Password
//	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError, "gagal memproses password")
//	}
//
//	_, err = con.Exec(
//		`INSERT INTO user (idRole, idStatusUser, username, email, nik, namaUser, noKontak, fotoKtp, password)
//   VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?`,
//		defaultIDRole, defaultIDStatusUser, req.Username, req.Email, tempValue, tempValue, tempValue, tempValue, string(hash),
//	)
//	if err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.JSON(http.StatusCreated, map[string]string{
//		"message": "registrasi berhasil",
//	})
//}
//
//func Profile(c *echo.Context) error {
//	id := c.FormValue("id")
//	namaUser := c.FormValue("namaUser")
//	nik := c.FormValue("nik")
//	noKontak := c.FormValue("noKontak")
//	fotoKtp := c.FormValue("fotoKtp")
//
//	conv_id, err := strconv.Atoi(id)
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//
//	result, err := models.Profile(conv_id, namaUser, nik, noKontak, fotoKtp)
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}
//
//func DeleteUser(c *echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//
//	result, err := models.DeleteUser(id)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}
