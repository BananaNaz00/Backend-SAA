package models

import (
	"myapp/db"
	"net/http"
)

type OpnameMaterial struct {
	Id                  int     `json:"id"`
	IdInventory         int     `json:"idInventory"`
	IdAlasanPenyesuaian int     `json:"idAlasanPenyesuaian"`
	StokSistem          float64 `json:"stokSistem"`
	StokFisik           float64 `json:"stokFisik"`
	Catatan             string  `json:"catatan"`
}

type InputOpnameMaterial struct {
	IdInventory         int     `json:"idInventory"`
	IdAlasanPenyesuaian int     `json:"idAlasanPenyesuaian"`
	StokSistem          float64 `json:"stokSistem"`
	StokFisik           float64 `json:"stokFisik"`
	Catatan             string  `json:"catatan"`
}

func DataOpnameMaterial() (Response, error) {
	var obj OpnameMaterial
	var arrobj []OpnameMaterial
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idOpname, idInventory, idAlasanPenyesuaian, stokSistem, stokFisik, catatan FROM opnamematerial"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdInventory, &obj.IdAlasanPenyesuaian, &obj.StokSistem, &obj.StokFisik, &obj.Catatan)
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

func SimpanOpnameMaterial(in InputOpnameMaterial) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO opnamematerial (idInventory, idAlasanPenyesuaian, stokSistem, stokFisik, catatan, createAt) VALUES (?, ?, ?, ?, ?, NOW())"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdInventory, in.IdAlasanPenyesuaian, in.StokSistem, in.StokFisik, in.Catatan)
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

func UpdateOpnameMaterial(id int, in InputOpnameMaterial) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE komoditas SET idInventory=?, idAlasanPenyesuaian=?, stokSistem=?, stokFisik=?, catatan=?, updateAt=Now() WHERE idOpname=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdInventory, in.IdAlasanPenyesuaian, in.StokSistem, in.StokFisik, in.Catatan, id)

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

func DeleteOpnameMaterial(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM opnamematerial WHERE idOpname=?"
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
