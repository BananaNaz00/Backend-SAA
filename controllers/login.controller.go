package controllers

//
//import (
//	"database/sql"
//	"errors"
//	"log"
//	"net/http"
//	"strings"
//
//	"github.com/labstack/echo/v5"
//	"golang.org/x/crypto/bcrypt"
//
//	"myapp/db"
//	"myapp/models"
//)
//
//func Register(c *echo.Context) error {
//	var req models.AuthRequest
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
//	// bcrypt hanya membaca maksimal 72 byte pertama password
//	if len(req.Password) < 6 || len(req.Password) > 72 {
//		return echo.NewHTTPError(http.StatusBadRequest, "password harus 6-72 karakter")
//	}
//
//	con := db.CreateCon()
//
//	var existingID int
//	err := con.QueryRow("SELECT idLogin FROM login WHERE username = ?", req.Username).Scan(&existingID)
//	if err == nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "username sudah digunakan")
//	}
//	//Untuk Hash Password
//	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return echo.NewHTTPError(http.StatusInternalServerError, "gagal memproses password")
//	}
//
//	_, err = con.Exec(
//		"INSERT INTO login (username, password) VALUES (?, ?)",
//		req.Username, string(hash),
//	)
//
//	return c.JSON(http.StatusCreated, map[string]string{
//		"message": "registrasi berhasil",
//	})
//}
//
//// Login mengecek username dan password terhadap data di tabel login.
//func Login(c *echo.Context) error {
//	var req models.AuthRequest
//	if err := c.Bind(&req); err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, "format request tidak valid")
//	}
//
//	var m models.Login
//	con := db.CreateCon()
//	err := con.QueryRow(
//		"SELECT idLogin, password FROM login WHERE username = ?",
//		strings.TrimSpace(req.Username),
//	).Scan(&m.IDLogin, &m.Password)
//
//	// Pesan sengaja sama supaya orang tidak bisa menebak username mana yang ada
//	invalid := echo.NewHTTPError(http.StatusUnauthorized, "username atau password salah")
//
//	if errors.Is(err, sql.ErrNoRows) {
//		return invalid
//	}
//	if err != nil {
//		log.Println("db error (login):", err)
//		return echo.NewHTTPError(http.StatusInternalServerError, "terjadi kesalahan server")
//	}
//
//	if bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(req.Password)) != nil {
//		return invalid
//	}
//
//	return c.JSON(http.StatusOK, map[string]any{
//		"message": "login berhasil",
//		"idLogin": m.IDLogin,
//	})
//}
