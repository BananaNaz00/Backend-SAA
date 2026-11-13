package models

import (
	"myapp/db"
	"net/http"
)

type KategoriInventory struct {
	Id           int    `json:"id"`
	NamaKategori string `json:"namaKategori"`
	Tipe         string `json:"tipe"`
}

func DataKategoriInventory() (Response, error) {
	var obj KategoriInventory
	var arrobj []KategoriInventory
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKategori, namaKategori, tipe FROM kategoriinventory"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaKategori, &obj.Tipe)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}
