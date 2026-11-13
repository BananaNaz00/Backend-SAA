package models

import (
	"myapp/db"
	"net/http"
)

type Komoditas struct {
	Id            int     `json:"id"`
	NamaKomoditas string  `json:"namaKomoditas"`
	EstimasiPanen int     `json:"estimasiPanen"`
	PhIdeal       float64 `json:"phIdeal"`
}

type InputKomoditas struct {
	NamaKomoditas string  `json:"namaKomoditas"`
	EstimasiPanen int     `json:"estimasiPanen"`
	PhIdeal       float64 `json:"phIdeal"`
}

func DataKomoditas() (Response, error) {
	var obj Komoditas
	var arrobj []Komoditas
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKomoditas, namaKomoditas, estimasiPanen, phIdeal FROM komoditas WHERE deleteAt IS NULL"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaKomoditas, &obj.EstimasiPanen, &obj.PhIdeal)
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

func SimpanKomoditas(in InputKomoditas) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO komoditas (namaKomoditas, estimasiPanen, phIdeal) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.NamaKomoditas, in.EstimasiPanen, in.PhIdeal)
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

func UpdateKomoditas(id int, in InputKomoditas) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE komoditas SET namaKomoditas=?, estimasiPanen=?, phIdeal=?, updateAt=Now() WHERE idKomoditas=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.NamaKomoditas, in.EstimasiPanen, in.PhIdeal, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}

func DeleteKomoditas(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE komoditas SET deleteAt=NOW() WHERE idKomoditas=? AND deleteAt IS NULL"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}
