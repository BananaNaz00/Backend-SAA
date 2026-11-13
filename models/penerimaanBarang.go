package models

import (
	"database/sql"
	"errors"
	"myapp/db"
	"net/http"
)

type PenerimaanBarang struct {
	Id                int     `json:"idPenerimaan"`
	IdInventory       int     `json:"idInventory"`
	TanggalPenerimaan string  `json:"tanggalPenerimaan"`
	NamaTokoVendor    string  `json:"namaTokoVendor"`
	Kuantitas         float64 `json:"kuantitas"`
	Catatan           string  `json:"catatan"`
}

type InputPenerimaanBarang struct {
	IdInventory       int     `json:"idInventory"`
	TanggalPenerimaan string  `json:"tanggalPenerimaan"`
	NamaTokoVendor    string  `json:"namaTokoVendor"`
	Kuantitas         float64 `json:"kuantitas"`
	Catatan           string  `json:"catatan"`
}

func DataPenerimaanBarang() (Response, error) {
	var obj PenerimaanBarang
	var arrobj []PenerimaanBarang
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT idPenerimaan, idInventory, tanggalPenerimaan, namaTokoVendor, kuantitas, catatan FROM penerimaanBarang"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdInventory, &obj.TanggalPenerimaan, &obj.NamaTokoVendor, &obj.Kuantitas, &obj.Catatan)
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

func SimpanPenerimaanBarang(in InputPenerimaanBarang) (Response, error) {
	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()
	if err != nil {
		return res, err
	}
	defer tx.Rollback()

	sqlStatement := "INSERT INTO penerimaanBarang (idInventory, tanggalPenerimaan, namaTokoVendor, kuantitas, catatan, createAt) VALUES (?, ?, ?, ?, ?, NOW())"
	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdInventory, in.TanggalPenerimaan, in.NamaTokoVendor, in.Kuantitas, in.Catatan)
	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	sqlUpdateInventory := "UPDATE inventory SET stokSistem=stokSistem+?, updateAt=NOW() WHERE idInventory=?"

	stmtUpdate, err := tx.Prepare(sqlUpdateInventory)
	if err != nil {
		return res, err
	}
	defer stmtUpdate.Close()

	_, err = stmtUpdate.Exec(in.Kuantitas, in.IdInventory)
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

func UpdatePenerimaanBarang(Id int, in InputPenerimaanBarang) (Response, error) {
	var res Response
	con := db.CreateCon()

	tx, err := con.Begin()
	if err != nil {
		return res, err
	}
	defer tx.Rollback()

	var stokAwal float64
	var idUpdate int
	sqlAwal := "SELECT kuantitas, idInventory FROM penerimaanbarang WHERE idPenerimaan=?"
	err = tx.QueryRow(sqlAwal, Id).Scan(&stokAwal, &idUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("inventory tidak ada")
		}
		return res, err
	}

	var updateStok = in.Kuantitas - stokAwal

	sqlStatement := "UPDATE penerimaanbarang SET tanggalPenerimaan=?, namaTokoVendor=?, kuantitas=?, catatan=?, updateAt=NOW() WHERE idPenerimaan=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.TanggalPenerimaan, in.NamaTokoVendor, in.Kuantitas, in.Catatan, Id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	sqlUpdateInventory := "UPDATE inventory SET stokSistem=stokSistem+?, updateAt=NOW() WHERE idInventory=?"

	stmtUpdate, err := tx.Prepare(sqlUpdateInventory)
	if err != nil {
		return res, err
	}
	defer stmtUpdate.Close()

	_, err = stmtUpdate.Exec(updateStok, idUpdate)
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
		"rows": rowsAffected,
	}

	return res, nil
}

func DeletePenerimaanBarang(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM penerimaanbarang WHERE idPenerimaan=?"
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
