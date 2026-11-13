package models

import (
	"myapp/db"
	"net/http"
)

type Role struct {
	Id            int    `json:"id"`
	NamaRole      string `json:"namaRole"`
	DeskripsiRole string `json:"deskripsiRole"`
}

type InputRole struct {
	NamaRole      string `json:"namaRole"`
	DeskripsiRole string `json:"deskripsiRole"`
}

func DataRole() (Response, error) {
	var obj Role
	var arrobj []Role
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idRole, namaRole, deskripsi FROM role"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.NamaRole, &obj.DeskripsiRole)
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

func SimpanRole(in InputRole) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO role (namaRole, deskripsi) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(in.NamaRole, in.DeskripsiRole)
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

func UpdateRole(id int, namaRole string, deskripsiRole string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE role SET namaRole=?, deskripsi=? WHERE idRole=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namaRole, deskripsiRole, id)

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

func DeleteRole(id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM role WHERE idRole=?"
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
