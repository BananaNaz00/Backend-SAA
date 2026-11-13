package models

import (
	"myapp/db"
	"net/http"
)

type ActivityLog struct {
	Id                  int    `json:"id"`
	IdKomoditasLahan    int    `json:"idKomoditasLahan"`
	IdKategoriAktivitas int    `json:"idKategoriAktivitas"`
	TanggalWaktu        string `json:"tanggalWaktu"`
	DeskripsiSingkat    string `json:"deskripsiSingkat"`
}

type InputActivityLog struct {
	IdKomoditasLahan    int    `json:"idKomoditasLahan"`
	IdKategoriAktivitas int    `json:"idKategoriAktivitas"`
	TanggalWaktu        string `json:"tanggalWaktu"`
	DeskripsiSingkat    string `json:"deskripsiSingkat"`
}

func DataActivityLog() (Response, error) {
	var obj ActivityLog
	var arrobj []ActivityLog
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idActivity, idKomoditasLahan, idKategoriAktivitas, tanggalWaktu, deskripsiSingkat FROM activitylog"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdKomoditasLahan, &obj.IdKategoriAktivitas, &obj.TanggalWaktu, &obj.DeskripsiSingkat)
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

func SimpanActivityLog(in InputActivityLog) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO activitylog (idKomoditasLahan, idKategoriAktivitas, tanggalWaktu, deskripsiSingkat, createAt) VALUES (?, ?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdKomoditasLahan, in.IdKategoriAktivitas, in.TanggalWaktu, in.DeskripsiSingkat)
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
