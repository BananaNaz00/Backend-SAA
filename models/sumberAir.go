package models

import (
	"myapp/db"
	"net/http"
)

type SumberAir struct {
	Id            int    `json:"id"`
	NamaSumberAir string `json:"namaSumberAir"`
}

func DataSumberAir() (Response, error) {
	var obj SumberAir
	var arrobj []SumberAir
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idSumberAir, namaSumberAir FROM sumberair"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaSumberAir)
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
