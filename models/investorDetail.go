package models

import (
	"myapp/db"
	"net/http"
)

type InvestorDetail struct {
	Id                  int    `json:"id"`
	IdUser              int    `json:"idUser"`
	NamaBank            string `json:"namaBank"`
	NomorRekening       string `json:"nomorRekening"`
	NamaPemilikRekening string `json:"namaPemilikRekening"`
	IdentitasLegalPDF   string `json:"identitasLegalPDF"`
	Npwp                string `json:"npwp"`
}

type InputInvestorDetail struct {
	IdUser              int    `json:"idUser"`
	NamaBank            string `json:"namaBank"`
	NomorRekening       string `json:"nomorRekening"`
	NamaPemilikRekening string `json:"namaPemilikRekening"`
	IdentitasLegalPDF   string `json:"identitasLegalPDF"`
	Npwp                string `json:"npwp"`
}

func DataInvestorDetail() (Response, error) {
	var obj InvestorDetail
	var arrobj []InvestorDetail
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT IdInvestor, idUser, namaBank, nomorRekening, namaPemilikRekening, identitasLegalPDF, npwp FROM investorDetail"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdUser, &obj.NamaBank, &obj.NomorRekening, &obj.NamaPemilikRekening, &obj.IdentitasLegalPDF, &obj.Npwp)
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

func SimpanInvestorDetail(in InputInvestorDetail) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO investordetail (idUser, namaBank, nomorRekening, namaPemilikRekening, identitasLegalPDF, npwp, createAt) VALUES (?, ?, ?, ?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdUser, in.NamaBank, in.NomorRekening, in.NamaPemilikRekening, in.IdentitasLegalPDF, in.Npwp)
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

func UpdateInvestorDetail(id int, in InputInvestorDetail) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE investordetail SET idUser=?, namaBank=?, nomorRekening=?, namaPemilikRekening=?, identitasLegalPDF=?, npwp=?, updateAt=Now() WHERE idInvestor=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdUser, in.NamaBank, in.NomorRekening, in.NamaPemilikRekening, in.IdentitasLegalPDF, in.Npwp, id)

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

func DeleteInvestorDetail(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM investordetail WHERE idInvestor=?"
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
