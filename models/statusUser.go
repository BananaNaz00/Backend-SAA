package models

import (
	"myapp/db"
	"net/http"
)

type StatusUser struct {
	Id             int    `json:"id"`
	NamaStatusUser string `json:"namaStatusUser"`
}

func DataStatusUser() (Response, error) {
	var obj StatusUser
	var arrobj []StatusUser
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idStatusUser, namaStatusUser FROM statususer"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaStatusUser)
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
