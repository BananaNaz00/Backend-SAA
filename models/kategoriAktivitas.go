package models

import (
	"myapp/db"
	"net/http"
)

type KategoriAktivitas struct {
	Id                    int    `json:"id"`
	NamaKategoriAktivitas string `json:"namaKategoriAktivitas"`
}

func DataKategoriAktivitas() (Response, error) {
	var obj KategoriAktivitas
	var arrobj []KategoriAktivitas
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKategoriAktivitas, namaKategoriAktivitas FROM kategoriaktivitas"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaKategoriAktivitas)
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
