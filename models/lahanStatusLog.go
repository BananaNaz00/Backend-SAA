package models

import (
	"myapp/db"
	"net/http"
)

type LahanStatusLog struct {
	Id               int    `json:"Id"`
	IdKomoditasLahan int    `json:"IdKomoditasLahan"`
	IdStatusLahan    int    `json:"IdStatusLahan"`
	TanggalPerubahan string `json:"TanggalPerubahan"`
	Catatan          string `json:"Catatan"`
}

type InputLahanStatusLog struct {
	IdKomoditasLahan int    `json:"IdKomoditasLahan"`
	IdStatusLahan    int    `json:"IdStatusLahan"`
	TanggalPerubahan string `json:"TanggalPerubahan"`
	Catatan          string `json:"Catatan"`
}

func DataLahanStatusLog() (Response, error) {
	var obj LahanStatusLog
	var arrobj []LahanStatusLog
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idStatusLog, idKomoditasLahan, idStatusLahan, tanggalPerubahan, catatan  FROM lahanstatuslog"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdKomoditasLahan, &obj.IdStatusLahan, &obj.TanggalPerubahan, &obj.Catatan)
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

func SimpanLahanStatusLog(in InputLahanStatusLog) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO lahanstatuslog (idKomoditasLahan, idStatusLahan, tanggalPerubahan, catatan, createAt) VALUES (?, ?, ?, ?, NOW())"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdKomoditasLahan, in.IdStatusLahan, in.TanggalPerubahan, in.Catatan)
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

//func UpdateLahanStatusLog(Id int, in InputLahanStatusLog) (Response, error) {
//	var res Response
//
//	con := db.CreateCon()
//
//	sqlStatement := "UPDATE lahanstatuslog SET idKomoditasLahan=?, idStatusLahan=?, tanggalPerubahan=?, catatan=? WHERE idStatusLog=?"
//	stmt, err := con.Prepare(sqlStatement)
//	if err != nil {
//		return res, err
//	}
//	defer stmt.Close()
//
//	result, err := stmt.Exec(in.IdKomoditasLahan, in.IdStatusLahan, in.TanggalPerubahan, in.Catatan, Id)
//
//	if err != nil {
//		return res, err
//	}
//
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		return res, err
//	}
//
//	res.Status = http.StatusOK
//	res.Message = "success"
//	res.Data = map[string]int64{
//		"rows": rowsAffected,
//	}
//
//	return res, nil
//}
