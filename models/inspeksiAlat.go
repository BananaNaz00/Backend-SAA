package models

import (
	"myapp/db"
	"net/http"
)

type InspeksiAlat struct {
	Id           int    `json:"id"`
	IdInventory  int    `json:"idInventory"`
	IdStatusAlat int    `json:"idStatusAlat"`
	Catatan      string `json:"catatan"`
}

type InputInspeksiAlat struct {
	IdInventory  int    `json:"idInventory"`
	IdStatusAlat int    `json:"idStatusAlat"`
	Catatan      string `json:"catatan"`
}

func DataInspeksiAlat() (Response, error) {
	var obj InspeksiAlat
	var arrobj []InspeksiAlat
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idInspeksiAlat, idInventory, idStatusAlat, Catatan FROM inspeksialat WHERE deleteAt IS NULL"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdInventory, &obj.IdStatusAlat, &obj.Catatan)
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

func SimpanInspeksiAlat(in InputInspeksiAlat) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO inspeksialat (idInventory, idStatusAlat, catatan, createAt) VALUES (?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdInventory, in.IdStatusAlat, in.Catatan)
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

func UpdateInspeksiAlat(id int, in InputInspeksiAlat) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE inspeksialat SET idInventory=?, idStatusAlat=?, catatan=?, updateAt=Now() WHERE idInspeksiAlat=?"
	stmt, err := con.Prepare(sqlStatement)
	defer stmt.Close()

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(in.IdInventory, in.IdStatusAlat, in.Catatan, id)

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

func DeleteInspeksiAlat(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE inspeksialat SET deleteAt=NOW() WHERE idInspeksiAlat=? AND deleteAt IS NULL"
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
