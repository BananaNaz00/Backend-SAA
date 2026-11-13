package models

import (
	"myapp/db"
	"net/http"
)

type KualitasSinyal struct {
	Id                 int    `json:"id"`
	NamaKualitasSinyal string `json:"namaKualitasSinyal"`
}

func DataKualitasSinyal() (Response, error) {
	var obj KualitasSinyal
	var arrobj []KualitasSinyal
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKualitasSinyal, namaKualitasSinyal FROM kualitassinyal"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaKualitasSinyal)
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
