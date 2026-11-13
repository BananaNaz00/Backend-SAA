package models

import (
	"myapp/db"
	"net/http"
)

type BlokLahan struct {
	Id            int    `json:"id"`
	IdLahan       int    `json:"idLahan"`
	NamaBlokLahan string `json:"namaBlokLahan"`
}

type InputBlokLahan struct {
	IdLahan       int    `json:"idLahan"`
	NamaBlokLahan string `json:"namaBlokLahan"`
}

func DataBlokLahan() (Response, error) {
	var obj BlokLahan
	var arrobj []BlokLahan
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idBlokLahan, idLahan, namaBlokLahan FROM bloklahan"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdLahan, &obj.NamaBlokLahan)
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

func SimpanBlokLahan(in InputBlokLahan) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO bloklahan (idLahan, namaBlokLahan) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.NamaBlokLahan)
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

func UpdateBlokLahan(Id int, in InputBlokLahan) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE bloklahan SET idLahan=?, namaBlokLahan=? WHERE idBlokLahan=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.NamaBlokLahan, Id)

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

func DeleteBlokLahan(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM blokLahan WHERE idBlokLahan=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

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
