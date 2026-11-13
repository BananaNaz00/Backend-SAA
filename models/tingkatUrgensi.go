package models

import (
	"myapp/db"
	"net/http"
)

type TingkatUrgensi struct {
	Id                 int    `json:"id"`
	NamaTingkatUrgensi string `json:"namaTingkatUrgensi"`
}

func DataTingkatUrgensi() (Response, error) {
	var obj TingkatUrgensi
	var arrobj []TingkatUrgensi
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idTingkatUrgensi, namaTingkatUrgensi FROM tingkaturgensi"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaTingkatUrgensi)
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

func SimpanTingkatUrgensi(namaTingkatUrgensi string) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO tingkaturgensi (namaTingkatUrgensi) VALUES (?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namaTingkatUrgensi)
	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"getIdLast": getIdLast,
	}

	return res, nil
}
