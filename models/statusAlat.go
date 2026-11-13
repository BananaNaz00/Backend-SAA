package models

import (
	"myapp/db"
	"net/http"
)

type StatusAlat struct {
	Id             int    `json:"id"`
	NamaStatusAlat string `json:"namaStatusAlat"`
}

func DataStatusAlat() (Response, error) {
	var obj StatusAlat
	var arrobj []StatusAlat
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idStatusAlat, namaStatusAlat FROM statusalat"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaStatusAlat)
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
