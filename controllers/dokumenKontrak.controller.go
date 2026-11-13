package controllers

import (
	"io"
	"myapp/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v5"
)

func DataDokumenKontrak(c *echo.Context) error {
	result, err := models.DataDokumenKontrak()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func GetDokumenKontrakByID(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := models.GetDokumenKontrakByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDokumenKontrak(c *echo.Context) error {
	idInvestStr := c.FormValue("idInvest")
	idInvest, _ := strconv.Atoi(idInvestStr)
	nomorKontrak := c.FormValue("nomorKontrak")
	namaDokumen := c.FormValue("namaDokumen")
	tanggalTandaTangan := c.FormValue("tanggalTandaTangan")
	tanggalMulai := c.FormValue("tanggalMulai")
	tanggalSelesai := c.FormValue("tanggalSelesai")

	fileDokumen, err := c.FormFile("fileDokumen")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	// untuk membuat folder jika tidak ada
	uploadDir := "Dokumen Kontrak/"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		_ = os.Mkdir(uploadDir, 0755)
	}

	var req models.InputDokumenKontrak
	req.IdInvest = idInvest
	req.NomorKontrak = nomorKontrak
	req.NamaDokumen = namaDokumen
	req.FileDokumen = ""
	req.TanggalTandaTangan = tanggalTandaTangan
	req.TanggalMulai = tanggalMulai
	req.TanggalSelesai = tanggalSelesai

	result, err := models.SimpanDokumenKontrak(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	dataMap := result.Data.(map[string]int64)
	getIdLast := dataMap["getIdLast"]
	idLast := strconv.Itoa(int(getIdLast))

	// membuka file
	src, err := fileDokumen.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	defer src.Close()

	namaFileBaru := idLast + "_" + fileDokumen.Filename
	// membuat file tujuan di folder yg sudah dibuat
	dstPath := uploadDir + namaFileBaru
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	defer dst.Close()

	// untuk copy isi file ke folder yg jdi tujuan
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = models.UpdateFileDokumenKontrak(getIdLast, namaFileBaru)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateDokumenKontrak(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	idInvestStr := c.FormValue("idInvest")
	idInvest, _ := strconv.Atoi(idInvestStr)
	nomorKontrak := c.FormValue("nomorKontrak")
	namaDokumen := c.FormValue("namaDokumen")
	tanggalTandaTangan := c.FormValue("tanggalTandaTangan")
	tanggalMulai := c.FormValue("tanggalMulai")
	tanggalSelesai := c.FormValue("tanggalSelesai")

	dataLama, err := models.GetDokumenKontrakByID(id)
	namaFileLama := ""
	if err == nil {
		namaFileLama = dataLama.FileDokumen
	}

	fileBaru, err := c.FormFile("fileDokumen")
	var namaFileFinal string
	uploadDir := "Dokumen Kontrak/"

	if err == nil {
		if namaFileLama != "" {
			pathFileLama := filepath.Join(uploadDir, namaFileLama)
			if _, err := os.Stat(pathFileLama); err == nil {
				os.Remove(pathFileLama)
			}
		}

		namaFileFinal = strconv.Itoa(id) + "_" + fileBaru.Filename
		dstPath := filepath.Join(uploadDir, namaFileFinal)

		src, err := fileBaru.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		defer src.Close()

		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			_ = os.Mkdir(uploadDir, 0755)
		}

		dst, err := os.Create(dstPath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

	} else {
		namaFileFinal = namaFileLama
	}

	var req models.InputDokumenKontrak
	req.IdInvest = idInvest
	req.NomorKontrak = nomorKontrak
	req.NamaDokumen = namaDokumen
	req.FileDokumen = namaFileFinal
	req.TanggalTandaTangan = tanggalTandaTangan
	req.TanggalMulai = tanggalMulai
	req.TanggalSelesai = tanggalSelesai

	result, err := models.UpdateDokumenKontrak(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteDokumenKontrak(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteDokumenKontrak(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
