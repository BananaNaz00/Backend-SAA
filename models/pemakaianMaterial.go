package models

import (
	"database/sql"
	"errors"
	"fmt"
	"myapp/db"
	"net/http"
)

type PemakaianMaterial struct {
	Id             int     `json:"id"`
	IdActivity     int     `json:"idActivity"`
	IdInventory    int     `json:"idInventory"`
	JumlahTerpakai float64 `json:"jumlahTerpakai"`
}

type InputPemakaianMaterial struct {
	IdActivity     int     `json:"idActivity"`
	IdInventory    int     `json:"idInventory"`
	JumlahTerpakai float64 `json:"jumlahTerpakai"`
}

func DataPemakaianMaterial() (Response, error) {
	var obj PemakaianMaterial
	var arrobj []PemakaianMaterial
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idPemakaian, idActivity, idInventory, jumlahTerpakai FROM pemakaianmaterial"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdActivity, &obj.IdInventory, &obj.JumlahTerpakai)
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

func SimpanPemakaianMaterial(in InputPemakaianMaterial) (Response, error) {

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()
	if err != nil {
		return res, err
	}
	defer tx.Rollback()

	var stokSistem float64
	sqlCekStok := "SELECT stokSistem FROM inventory WHERE idInventory=?"
	err = tx.QueryRow(sqlCekStok, in.IdInventory).Scan(&stokSistem)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("inventory tidak ada")
		}
		return res, err
	}

	if stokSistem < in.JumlahTerpakai {
		return res, fmt.Errorf("stok tidak mencukupi")
	}

	sqlStatement := "INSERT INTO pemakaianmaterial (idActivity, idInventory, jumlahTerpakai) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(in.IdActivity, in.IdInventory, in.JumlahTerpakai)
	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	sqlUpdateInventory := "UPDATE inventory SET stokSistem=stokSistem-?, updateAt=NOW() WHERE idInventory=?"
	stmtUpdate, err := tx.Prepare(sqlUpdateInventory)
	if err != nil {
		return res, err
	}
	defer stmtUpdate.Close()

	_, err = stmtUpdate.Exec(in.JumlahTerpakai, in.IdInventory)
	if err != nil {
		return res, err
	}

	err = tx.Commit()
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

func UpdatePemakaianMaterial(id int, in InputPemakaianMaterial) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pemakaianmaterial SET idActivity=?, idInventory=?, jumlahTerpakai=? WHERE idPemakaian=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdActivity, in.IdInventory, in.JumlahTerpakai, id)

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

func DeletePemakaianMaterial(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pemakaianmaterial WHERE idPemakaian=?"
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
