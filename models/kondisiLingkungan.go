package models

import (
	"myapp/db"
	"net/http"
)

type KondisiLingkungan struct {
	Id              int     `json:"Id"`
	IdLahan         int     `json:"IdLahan"`
	NilaiPh         float64 `json:"NilaiPh"`
	SuhuTanah       float64 `json:"SuhuTanah"`
	KelembapanTanah float64 `json:"KelembapanTanah"`
	Tanggal         string  `json:"Tanggal"`
}

type InputKondisiLingkungan struct {
	IdLahan         int     `json:"IdLahan"`
	NilaiPh         float64 `json:"NilaiPh"`
	SuhuTanah       float64 `json:"SuhuTanah"`
	KelembapanTanah float64 `json:"KelembapanTanah"`
	Tanggal         string  `json:"Tanggal"`
}

func DataKondisiLingkungan() (Response, error) {
	var obj KondisiLingkungan
	var arrobj []KondisiLingkungan
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idKondisi, idLahan, nilaiPh, suhuTanah, kelembapanTanah, tanggal FROM kondisilingkungan"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdLahan, &obj.NilaiPh, &obj.SuhuTanah, &obj.KelembapanTanah, &obj.Tanggal)
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

func SimpanKondisiLingkungan(in InputKondisiLingkungan) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO kondisilingkungan (idLahan, nilaiPh, suhuTanah, kelembapanTanah, tanggal) VALUES (?, ?, ?, ?, ?)"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.NilaiPh, in.SuhuTanah, in.KelembapanTanah, in.Tanggal)
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

func UpdateKondisiLingkungan(Id int, in InputKondisiLingkungan) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE kondisilingkungan SET idLahan=?, nilaiPh=?, suhuTanah=?, kelembapanTanah=?, tanggal=? WHERE idKondisi=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.NilaiPh, in.SuhuTanah, in.KelembapanTanah, in.Tanggal, Id)

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

func DeleteKondisiLingkungan(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM kondisilingkungan WHERE idKondisi=?"
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
