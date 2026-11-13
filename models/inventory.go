package models

import (
	"myapp/db"
	"net/http"
)

type Inventory struct {
	Id         int     `json:"id"`
	IdKategori int     `json:"idKategori"`
	IdLahan    int     `json:"idLahan"`
	IdSatuan   int     `json:"idSatuan"`
	NamaBarang string  `json:"namaBarang"`
	StokSistem float64 `json:"stokSistem"`
}

type InputInventory struct {
	IdKategori int     `json:"idKategori"`
	IdLahan    int     `json:"idLahan"`
	IdSatuan   int     `json:"idSatuan"`
	NamaBarang string  `json:"namaBarang"`
	StokSistem float64 `json:"stokSistem"`
}

func DataInventory() (Response, error) {
	var obj Inventory
	var arrobj []Inventory
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idInventory, idKategori, idLahan, idSatuan, namaBarang, stokSistem FROM inventory WHERE deleteAt IS NULL"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdKategori, &obj.IdLahan, &obj.IdSatuan, &obj.NamaBarang, &obj.StokSistem)
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

func SimpanInventory(in InputInventory) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO inventory (idKategori, idLahan, idSatuan, namaBarang, stokSistem, createAt) VALUES (?, ?, ?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdKategori, in.IdLahan, in.IdSatuan, in.NamaBarang, in.StokSistem)
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

func UpdateInventory(id int, in InputInventory) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE inventory SET idKategori=?, idLahan=?, idSatuan=?, namaBarang=?, stokSistem=?, updateAt=NOW() WHERE idInventory=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdKategori, in.IdLahan, in.IdSatuan, in.NamaBarang, in.StokSistem, id)

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

func DeleteInventory(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE inventory SET deleteAt=NOW() WHERE idInventory=? AND deleteAt IS NULL"
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
