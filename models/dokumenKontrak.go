package models

import (
	"myapp/db"
	"net/http"
)

type DokumenKontrak struct {
	Id                 int    `json:"id"`
	IdInvest           int    `json:"IdInvest"`
	NomorKontrak       string `json:"NomorKontrak"`
	NamaDokumen        string `json:"namaDokumen"`
	FileDokumen        string `json:"fileDokumen"`
	TanggalTandaTangan string `json:"tanggalTandaTangan"`
	TanggalMulai       string `json:"tanggalMulai"`
	TanggalSelesai     string `json:"tanggalSelesai"`
}

type InputDokumenKontrak struct {
	IdInvest           int    `json:"IdInvest"`
	NomorKontrak       string `json:"NomorKontrak"`
	NamaDokumen        string `json:"namaDokumen"`
	FileDokumen        string `json:"fileDokumen"`
	TanggalTandaTangan string `json:"tanggalTandaTangan"`
	TanggalMulai       string `json:"tanggalMulai"`
	TanggalSelesai     string `json:"tanggalSelesai"`
}

func DataDokumenKontrak() (Response, error) {
	var obj DokumenKontrak
	var arrobj []DokumenKontrak
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idDokumenKontrak, idInvest, nomorKontrak, namaDokumen, fileDokumen, tanggalTandaTangan, tanggalMulai, tanggalSelesai FROM dokumenKontrak WHERE deleteAt IS NULL"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdInvest, &obj.NomorKontrak, &obj.NamaDokumen, &obj.FileDokumen, &obj.TanggalTandaTangan, &obj.TanggalMulai, &obj.TanggalSelesai)
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

func GetDokumenKontrakByID(id int) (InputDokumenKontrak, error) {
	var obj InputDokumenKontrak
	con := db.CreateCon()

	sqlStatement := "SELECT idInvest, nomorKontrak, namaDokumen, fileDokumen, tanggalTandaTangan, tanggalMulai, tanggalSelesai FROM dokumenkontrak WHERE idDokumenKontrak=?"

	err := con.QueryRow(sqlStatement, id).Scan(&obj.IdInvest, &obj.NomorKontrak, &obj.NamaDokumen, &obj.FileDokumen, &obj.TanggalTandaTangan, &obj.TanggalMulai, &obj.TanggalSelesai)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func SimpanDokumenKontrak(in InputDokumenKontrak) (Response, error) {

	var res Response

	temp := ""
	con := db.CreateCon()
	sqlStatement := "INSERT INTO dokumenkontrak (idInvest, nomorKontrak, namaDokumen, fileDokumen, tanggalTandaTangan, tanggalMulai, tanggalSelesai, createAt) VALUES (?, ?, ?,?, ?, ?, ?, NOW())"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(in.IdInvest, in.NomorKontrak, in.NamaDokumen, temp, in.TanggalTandaTangan, in.TanggalMulai, in.TanggalSelesai)
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

func UpdateFileDokumenKontrak(id int64, namaFile string) error {
	con := db.CreateCon()

	sqlStatement := "UPDATE dokumenkontrak SET fileDokumen=? WHERE idDokumenKontrak=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(namaFile, id)
	return err
}

func UpdateDokumenKontrak(Id int, in InputDokumenKontrak) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE dokumenkontrak SET idInvest=?, nomorKontrak=?, namaDokumen=?, fileDokumen=?, tanggalTandaTangan=?, tanggalMulai=?, tanggalSelesai=?, updateAt=NOW() WHERE idDokumenKontrak=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdInvest, in.NomorKontrak, in.NamaDokumen, in.FileDokumen, in.TanggalTandaTangan, in.TanggalMulai, in.TanggalSelesai, Id)

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

func DeleteDokumenKontrak(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE dokumenkontrak SET deleteAt=NOW() WHERE idDokumenKontrak=? AND deleteAt IS NULL"
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
