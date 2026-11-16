package models

import (
	"database/sql"
	"myapp/db"
	"net/http"
)

type Lahan struct {
	Id                int             `json:"Id"`
	IdAdmin           int             `json:"IdAdmin"`
	IdInvestor        int             `json:"IdInvestor"`
	IdPenanggungJawab int             `json:"IdPenanggungJawab"`
	IdStatusLahan     int             `json:"IdStatusLahan"`
	NamaLahan         string          `json:"NamaLahan"`
	Koordinat         string          `json:"Koordinat"`
	Alamat            string          `json:"Alamat"`
	AreaPolygon       string          `json:"AreaPolygon"`
	DokumenSertifikat string          `json:"DokumenSertifikat"`
	LuasLahan         float64         `json:"LuasLahan"`
	NilaiLahan        float64         `json:"NilaiLahan"`
	CreateBy          string          `json:"CreateBy"`
	UpdateBy          string          `json:"UpdateBy"`
	Deskripsi         string          `json:"Deskripsi"`
	Komoditas         []komoditasItem `json:"Komoditas"`
}

type komoditasItem struct {
	Komoditas           string  `json:"Komoditas"`
	PersentaseKomoditas float64 `json:"PersentaseKomoditas"`
}

type InputLahan struct {
	IdAdmin           int             `json:"IdAdmin"`
	IdInvestor        int             `json:"IdInvestor"`
	IdPenanggungJawab int             `json:"IdPenanggungJawab"`
	IdStatusLahan     int             `json:"IdStatusLahan"`
	NamaLahan         string          `json:"NamaLahan"`
	Koordinat         string          `json:"Koordinat"`
	Alamat            string          `json:"Alamat"`
	AreaPolygon       string          `json:"AreaPolygon"`
	DokumenSertifikat string          `json:"DokumenSertifikat"`
	LuasLahan         float64         `json:"LuasLahan"`
	NilaiLahan        float64         `json:"NilaiLahan"`
	CreateBy          string          `json:"CreateBy"`
	UpdateBy          string          `json:"UpdateBy"`
	Deskripsi         string          `json:"Deskripsi"`
	Komoditas         []komoditasItem `json:"Komoditas"`
}

func getKomoditasByLahanId(con *sql.DB, idLahan int) ([]komoditasItem, error) {
	var list []komoditasItem

	sqlStatement := `SELECT k.namaKomoditas, kl.persentaseArea FROM komoditaslahan kl JOIN bloklahan bl ON kl.idBlokLahan = bl.idBlokLahan JOIN komoditas k ON kl.idKomoditas = k.idKomoditas WHERE bl.idLahan = ?`

	rows, err := con.Query(sqlStatement, idLahan)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var k komoditasItem
		if err := rows.Scan(&k.Komoditas, &k.PersentaseKomoditas); err != nil {
			return list, err
		}
		list = append(list, k)
	}
	return list, nil
}

func DataLahan() (Response, error) {
	var obj Lahan
	var arrobj []Lahan
	var res Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatment := "SELECT idLahan, idAdmin, idInvestor, idPenanggungJawab, idStatusLahan, namaLahan, koordinat, alamat, areaPolygon, dokumenSertifikat, luasLahan, nilaiLahan, createBy, updateBy, deskripsi FROM lahan WHERE deleteAt IS NULL"

	rows, err := con.Query(sqlStatment)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.IdAdmin, &obj.IdInvestor, &obj.IdPenanggungJawab, &obj.IdStatusLahan, &obj.NamaLahan, &obj.Koordinat, &obj.Alamat, &obj.AreaPolygon, &obj.DokumenSertifikat, &obj.LuasLahan, &obj.NilaiLahan, &obj.CreateBy, &obj.UpdateBy, &obj.Deskripsi)
		if err != nil {
			return res, err
		}

		komoditasList, err := getKomoditasByLahanId(con, obj.Id)
		if err == nil {
			obj.Komoditas = komoditasList
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}

func GetLahanById(id int) (Lahan, error) {
	var res Response
	var obj Lahan
	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "SELECT idLahan, idAdmin, idInvestor, idPenanggungJawab, idStatusLahan, namaLahan, koordinat, alamat, areaPolygon, dokumenSertifikat, luasLahan, nilaiLahan, createBy, updateBy, deskripsi FROM lahan WHERE idLahan=?"

	err := con.QueryRow(sqlStatement, id).Scan(&obj.Id, &obj.IdAdmin, &obj.IdInvestor, &obj.IdPenanggungJawab, &obj.IdStatusLahan, &obj.NamaLahan, &obj.Koordinat, &obj.Alamat, &obj.AreaPolygon, &obj.DokumenSertifikat, &obj.LuasLahan, &obj.NilaiLahan, &obj.CreateBy, &obj.UpdateBy, &obj.Deskripsi)
	if err != nil {
		return obj, err
	}

	komoditasList, err := getKomoditasByLahanId(con, obj.Id)
	if err == nil {
		obj.Komoditas = komoditasList
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = obj

	return obj, nil
}

func SimpanLahan(in InputLahan) (Response, error) {
	var res Response
	con := db.CreateCon()
	defer con.Close()
	sqlStatement := `INSERT INTO lahan (idAdmin, idInvestor, idPenanggungJawab, idStatusLahan, namaLahan, koordinat, alamat, areaPolygon, dokumenSertifikat, luasLahan, nilaiLahan, createBy, deskripsi, createAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		in.IdAdmin, in.IdInvestor, in.IdPenanggungJawab, in.IdStatusLahan,
		in.NamaLahan, in.Koordinat, in.Alamat, in.AreaPolygon, in.DokumenSertifikat,
		in.LuasLahan, in.NilaiLahan, in.CreateBy, in.Deskripsi,
	)
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

func UpdateLahan(Id int, in InputLahan) (Response, error) {
	var res Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE lahan SET idAdmin=?, idInvestor=?, idPenanggungJawab=?, idStatusLahan=?, namaLahan=?, koordinat=?, alamat=?, areaPolygon=?, dokumenSertifikat=?, luasLahan=?, nilaiLahan=?, updateBy=?, deskripsi=?, updateAt=NOW() WHERE idLahan=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(in.IdAdmin, in.IdInvestor, in.IdPenanggungJawab, in.IdStatusLahan, in.NamaLahan, in.Koordinat, in.Alamat, in.AreaPolygon, in.DokumenSertifikat, in.LuasLahan, in.NilaiLahan, in.UpdateBy, in.Deskripsi, Id)

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

func DeleteLahan(id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE lahan SET deleteAt=NOW() WHERE idLahan=? AND deleteAt IS NULL"
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
