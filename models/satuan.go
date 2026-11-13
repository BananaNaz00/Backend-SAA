package models

import (
	"myapp/db"
	"net/http"
)

type Satuan struct {
	Id         int    `json:"id"`
	NamaSatuan string `json:"namaSatuan"`
}

func Asatuan() (Response, error) {
	var obj Satuan
	var arrobj []Satuan
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT idSatuan, namaSatuan FROM Satuan"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&obj.Id, &obj.NamaSatuan)
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
