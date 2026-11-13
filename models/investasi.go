package models

import (
	"myapp/db"
	"net/http"
)

type Investasi struct {
	Id                int     `json:"Id"`
	IdStatusInvestasi int     `json:"IdStatusInvestasi"`
	IdLahan           int     `json:"IdLahan"`
	Nominal           float64 `json:"Nominal"`
	MetodePembayaran  string  `json:"MetodePembayaran"`
	BuktiPembayaran   string  `json:"BuktiPembayaran"`
	TanggalInvestasi  string  `json:"TanggalInvestasi"`
}

type InputInvestasi struct {
	IdStatusInvestasi int     `json:"IdStatusInvestasi"`
	IdLahan           int     `json:"IdLahan"`
	Nominal           float64 `json:"Nominal"`
	MetodePembayaran  string  `json:"MetodePembayaran"`
	BuktiPembayaran   string  `json:"BuktiPembayaran"`
	TanggalInvestasi  string  `json:"TanggalInvestasi"`
}

func DataInvestasi() (Response, error) {
	var obj Investasi
	var arrobj []Investasi
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idInvest, idStatusInvestasi, idLahan, nominal, metodePembayaran, buktiPembayaran, tanggalInvestasi FROM investasi WHERE deleteAt IS NULL"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdStatusInvestasi, &obj.IdLahan, &obj.Nominal, &obj.MetodePembayaran, &obj.BuktiPembayaran, &obj.TanggalInvestasi)
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

func SimpanInvestasi(in InputInvestasi) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO investasi (idStatusInvestasi, idLahan, nominal, metodePembayaran, buktiPembayaran, tanggalInvestasi, createAt) VALUES (?, ?, ?, ?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdStatusInvestasi, in.IdLahan, in.Nominal, in.MetodePembayaran, in.BuktiPembayaran, in.TanggalInvestasi)
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

func UpdateInvestasi(Id int, in InputInvestasi) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE investasi SET idStatusInvestasi=?, idLahan=?, nominal=?, metodePembayaran=?, buktiPembayaran=?, tanggalInvestasi=?, updateAt=NOW() WHERE idInvest=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdStatusInvestasi, in.IdLahan, in.Nominal, in.MetodePembayaran, in.BuktiPembayaran, in.TanggalInvestasi, Id)

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

func DeleteInvestasi(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE investasi SET deleteAt=NOW() WHERE idInvest=? AND deleteAt IS NULL"
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
