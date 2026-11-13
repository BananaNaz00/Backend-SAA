package models

import (
	"myapp/db"
	"net/http"
)

type KategoriKendala struct {
	Id                  int    `json:"id"`
	NamaKategoriKendala string `json:"namaKategoriKendala"`
}

func DataKategoriKendala() (Response, error) {
	var obj KategoriKendala
	var arrobj []KategoriKendala
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKategoriKendala, namaKategoriKendala FROM kategorikendala"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaKategoriKendala)
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
