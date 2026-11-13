package models

import (
	"myapp/db"
	"net/http"
)

type StatusInvestasi struct {
	Id                  int    `json:"id"`
	NamaStatusInvestasi string `json:"namaStatusInvestasi"`
}

func DataStatusInvestasi() (Response, error) {
	var obj StatusInvestasi
	var arrobj []StatusInvestasi
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idStatusInvestasi, namaStatusInvestasi FROM statusinvestasi"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaStatusInvestasi)
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
