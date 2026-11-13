package models

import (
	"myapp/db"
	"net/http"
)

type Insiden struct {
	Id                   int    `json:"id"`
	IdLahan              int    `json:"idLahan"`
	IdKategoriKendala    int    `json:"idKategoriKendala"`
	IdTingkatUrgensi     int    `json:"idTingkatUrgensi"`
	IdStatusPenyelesaian int    `json:"idStatusPenyelesaian"`
	DeskripsiKendala     string `json:"deskripsiKendala"`
	TanggalLapor         string `json:"tanggalLapor"`
}

type InputInsiden struct {
	IdLahan              int    `json:"idLahan"`
	IdKategoriKendala    int    `json:"idKategoriKendala"`
	IdTingkatUrgensi     int    `json:"idTingkatUrgensi"`
	IdStatusPenyelesaian int    `json:"idStatusPenyelesaian"`
	DeskripsiKendala     string `json:"deskripsiKendala"`
	TanggalLapor         string `json:"tanggalLapor"`
}

func DataInsiden() (Response, error) {
	var obj Insiden
	var arrobj []Insiden
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idInsiden, idLahan, idKategoriKendala, idTingkatUrgensi, idStatusPenyelesaian, deskripsiKendala, tanggalLapor FROM insiden"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdLahan, &obj.IdKategoriKendala, &obj.IdTingkatUrgensi, &obj.IdStatusPenyelesaian, &obj.DeskripsiKendala, &obj.TanggalLapor)
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

func SimpanInsiden(in InputInsiden) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO insiden (idLahan, idKategoriKendala, idTingkatUrgensi, idStatusPenyelesaian, deskripsiKendala, tanggalLapor, createAt) VALUES (?, ?, ?, ?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.IdKategoriKendala, in.IdTingkatUrgensi, in.IdStatusPenyelesaian, in.DeskripsiKendala, in.TanggalLapor)
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

func UpdateInsiden(id int, in InputInsiden) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE insiden SET idLahan=?, idKategoriKendala=?, idTingkatUrgensi=?, idStatusPenyelesaian=?, deskripsiKendala=?, tanggalLapor=?, updateAt=Now() WHERE idInsiden=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.IdKategoriKendala, in.IdTingkatUrgensi, in.IdStatusPenyelesaian, in.DeskripsiKendala, in.TanggalLapor, id)

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

func DeleteInsiden(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM insiden WHERE idInsiden=?"
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
