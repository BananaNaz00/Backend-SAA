package models

import (
	"myapp/db"
	"net/http"
)

type StatusLahan struct {
	Id              int    `json:"id"`
	NamaStatusLahan string `json:"namaStatusLahan"`
}

func DataStatusLahan() (Response, error) {
	var obj StatusLahan
	var arrobj []StatusLahan
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idStatusLahan, namaStatusLahan FROM statuslahan"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaStatusLahan)
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
