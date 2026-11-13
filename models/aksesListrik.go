package models

import (
	"myapp/db"
	"net/http"
)

type AksesListrik struct {
	Id               int    `json:"id"`
	NamaAksesListrik string `json:"namaAksesListrik"`
}

func DataAksesListrik() (Response, error) {
	var obj AksesListrik
	var arrobj []AksesListrik
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idAksesListrik, namaAksesListrik FROM akseslistrik"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaAksesListrik)
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
