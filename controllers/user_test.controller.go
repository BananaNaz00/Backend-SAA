package controllers

import (
	"database/sql"
	"io"
	"myapp/db"
	"myapp/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func DataUser(c *echo.Context) error {
	result, err := models.DataUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func Register(c *echo.Context) error {
	var input models.RegisterRequest
	temp := ""
	defaultStatus := 1

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validasi gagal",
			"error":   err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "gagal2"})
	}

	user := models.User{
		IdRole:   input.IdRole,
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	con := db.CreateCon()
	sqlStatement := "INSERT INTO user (idRole, idStatusUser, username, email, nik, password, namaUser, noKontak, fotoKTP, createAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "gagal3"})
	}
	_, err = stmt.Exec(user.IdRole, defaultStatus, user.Username, user.Email, temp, string(hashedPassword), temp, temp, temp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Register Berhasil!!!",
		"user": map[string]interface{}{
			"idRole":   user.IdRole,
			"username": user.Username,
			"email":    user.Email,
			"password": user.Password},
	})
}

func Login(c *echo.Context) error {
	var input models.Login

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := validate.Struct(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validasi gagal",
			"error":   err.Error(),
		})
	}

	//Check Username
	con := db.CreateCon()
	//defer con.Close()

	var existingUsername int
	var hashedPassword string
	err := con.QueryRow("SELECT idUser, password FROM user WHERE username = ?", input.Username).Scan(&existingUsername, &hashedPassword)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "username tidak terdaftar",
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	//Check Password
	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password))

	if errMatchPassword != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Password salah",
		})
	}

	//Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUsername,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenstring, err := token.SignedString([]byte(os.Getenv("jwt_secret_rahasia")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal Membuat Token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Berhasil!!",
		"token":   tokenstring,
	})
}

func UpdateProfile(c *echo.Context) error {
	idStr := c.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "ID tidak valid"})
	}
	nik := c.FormValue("nik")
	namaUser := c.FormValue("namaUser")
	noKontak := c.FormValue("noKontak")

	DataLama, err := models.GetUserById(id) // Sesuaikan tipe data parameter di model jika butuh int/string
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User tidak ditemukan"})
	}

	fileFotoKtp, errFile := c.FormFile("fotoKTP")
	if errFile != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// untuk membuat folder jika tidak ada
	uploadDir := "Foto Ktp User/"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		_ = os.Mkdir(uploadDir, 0755)
	}

	var pathFoto string
	var namaFotoUser string

	if err == nil {
		src, err := fileFotoKtp.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Gagal Membuka File",
			})
		}
		defer src.Close()

		ext := filepath.Ext(fileFotoKtp.Filename)

		namaFotoUser = DataLama.Username + ext
		pathFoto = filepath.Join(uploadDir, namaFotoUser)

		dst, err := os.Create(pathFoto)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Gagal Menyimpan File",
			})
		}
		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Gagal Menyalin File",
			})
		}

	} else {
		pathFoto = DataLama.FotoKTP
	}

	profile := models.Profile{
		Nik:      nik,
		NamaUser: namaUser,
		NoKontak: noKontak,
		FotoKTP:  namaFotoUser,
	}
	// 3. Lanjutkan proses update data teks & path foto ke database MariaDB kamu di sini...
	result, err := models.EditProfile(id, profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal memperbarui ke database: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
