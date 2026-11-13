package models

import (
	"myapp/db"
	"net/http"
)

type StatusPenyelesaian struct {
	Id                     int    `json:"id"`
	NamaStatusPenyelesaian string `json:"namaStatusPenyelesaian"`
}

func DataStatusPenyelesaian() (Response, error) {
	var obj StatusPenyelesaian
	var arrobj []StatusPenyelesaian
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idStatusPenyelesaian, namaStatusPenyelesaian FROM statuspenyelesaian"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaStatusPenyelesaian)
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
