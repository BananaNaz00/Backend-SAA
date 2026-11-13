package models

import (
	"myapp/db"
	"net/http"
)

type AksesJalan struct {
	Id             int    `json:"id"`
	NamaAksesJalan string `json:"namaAksesJalan"`
}

func DataAksesJalan() (Response, error) {
	var obj AksesJalan
	var arrobj []AksesJalan
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idAksesJalan, namaAksesJalan FROM aksesjalan"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaAksesJalan)
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
