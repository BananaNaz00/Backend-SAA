package models

import (
	"myapp/db"
	"net/http"
)

type AlasanPenyesuaian struct {
	Id                    int    `json:"id"`
	NamaAlasanPenyesuaian string `json:"namaAlasanPenyesuaian"`
}

func DataAlasanPenyesuaian() (Response, error) {
	var obj AlasanPenyesuaian
	var arrobj []AlasanPenyesuaian
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idAlasanPenyesuaian, namaAlasanPenyesuaian FROM alasanpenyesuaian"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaAlasanPenyesuaian)
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
