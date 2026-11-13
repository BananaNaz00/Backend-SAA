package models

import (
	"myapp/db"
	"net/http"
)

type KomoditasLahan struct {
	Id             int     `json:"id"`
	IdBlokLahan    int     `json:"idBlokLahan"`
	IdKomoditas    int     `json:"idKomoditas"`
	PersentaseArea float64 `json:"persentaseArea"`
}

type InputKomoditasLahan struct {
	IdBlokLahan    int     `json:"idBlokLahan"`
	IdKomoditas    int     `json:"idKomoditas"`
	PersentaseArea float64 `json:"persentaseArea"`
}

func DataKomoditasLahan() (Response, error) {
	var obj KomoditasLahan
	var arrobj []KomoditasLahan
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKomoditasLahan, idBlokLahan, idKomoditas, persentaseArea FROM komoditaslahan"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdBlokLahan, &obj.IdKomoditas, &obj.PersentaseArea)
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

func SimpanKomoditasLahan(in InputKomoditasLahan) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO komoditaslahan (idBlokLahan, idKomoditas, persentaseArea) VALUES (?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdBlokLahan, in.IdKomoditas, in.PersentaseArea)
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

func UpdateKomoditasLahan(id int, in InputKomoditasLahan) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE komoditaslahan SET idBlokLahan=?, idKomoditas=?, persentaseArea=? WHERE idKomoditasLahan=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdBlokLahan, in.IdKomoditas, in.PersentaseArea, id)

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

func DeleteKomoditasLahan(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM komoditaslahan WHERE idKomoditasLahan=?"
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
