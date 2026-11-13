package models

import (
	"myapp/db"
	"net/http"
)

type FasilitasLahan struct {
	Id               int `json:"id"`
	IdLahan          int `json:"idLahan"`
	IdSumberAir      int `json:"idSumberAir"`
	IdAksesListrik   int `json:"idAksesListrik"`
	IdAksesJalan     int `json:"idAksesJalan"`
	IdKualitasSinyal int `json:"idKualitasSinyal"`
}

// Dipakai controller untuk Bind body JSON saat create/update.
type InputFasilitasLahan struct {
	IdLahan          int `json:"idLahan"`
	IdSumberAir      int `json:"idSumberAir"`
	IdAksesListrik   int `json:"idAksesListrik"`
	IdAksesJalan     int `json:"idAksesJalan"`
	IdKualitasSinyal int `json:"idKualitasSinyal"`
}

func DataFasilitasLahan() (Response, error) {
	var obj FasilitasLahan
	var arrobj []FasilitasLahan
	var res Response

	con := db.CreateCon()

	sqlStatment := "SELECT idFasilitas, idLahan, idSumberAir, idAksesListrik, idAksesJalan, idKualitasSinyal FROM fasilitaslahan"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdLahan, &obj.IdSumberAir, &obj.IdAksesListrik, &obj.IdAksesJalan, &obj.IdKualitasSinyal)
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

func SimpanFasilitasLahan(in InputFasilitasLahan) (Response, error) {

	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO fasilitaslahan (idLahan, idSumberAir, idAksesListrik, idAksesJalan, idkualitasSinyal, createAt) VALUES (?, ?, ?, ?, ?, NOW())"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.IdSumberAir, in.IdAksesListrik, in.IdAksesJalan, in.IdKualitasSinyal)
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

func UpdateFasilitasLahan(id int, in InputFasilitasLahan) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE fasilitaslahan SET idLahan=?, idSumberAir=?, idAksesListrik=?, idAksesJalan=?, idkualitasSinyal=?, updateAt=Now() WHERE idFasilitas=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdLahan, in.IdSumberAir, in.IdAksesListrik, in.IdAksesJalan, in.IdKualitasSinyal, id)

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

func DeleteFasilitasLahan(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM fasilitaslahan WHERE idFasilitas=?"
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
