package models

//
//import (
//	"myapp/db"
//	"net/http"
//)
//
//type User struct {
//	IdUser       int    `json:"idUser"`
//	IdRole       int    `json:"idRole"`
//	IdStatusUser int    `json:"idStatusUser"`
//	Username     string `json:"username"`
//	Email        string `json:"email"`
//	Nik          string `json:"nik"`
//	Password     string `json:"-"`
//	NamaUser     string `json:"namaUser"`
//	NoKontak     string `json:"noKontak"`
//	FotoKTP      string `json:"fotoKTP"`
//}
//
//type RegisterRequest struct {
//	Username string `json:"username"`
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}
//
//func DataUser() (Response, error) {
//	var obj User
//	var arrobj []User
//	var res Response
//
//	con := db.CreateCon()
//
//	sqlStatment := "SELECT idUser, idRole, idStatusUser, username, email, nik, namaUser, noKontak, fotoKTP FROM user WHERE deleteAt IS NULL"
//
//	rows, err := con.Query(sqlStatment)
//	if err != nil {
//		return res, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		err = rows.Scan(&obj.IdUser, &obj.IdRole, &obj.IdStatusUser, &obj.Username, &obj.Email, &obj.Nik, &obj.NamaUser, &obj.NoKontak, &obj.FotoKTP)
//		if err != nil {
//			return res, err
//		}
//		arrobj = append(arrobj, obj)
//	}
//
//	res.Status = http.StatusOK
//	res.Message = "success"
//	res.Data = arrobj
//
//	return res, nil
//}
//
//func Profile(id int, namaUser string, nik string, noKontak string, fotoKtp string) (Response, error) {
//	var res Response
//
//	con := db.CreateCon()
//
//	sqlStatement := "UPDATE user SET namaUser=?, nik=?, noKontak=?, fotoKtp=?, createAt=NOW() WHERE idUser=?"
//	stmt, err := con.Prepare(sqlStatement)
//
//	if err != nil {
//		return res, err
//	}
//
//	result, err := stmt.Exec(namaUser, nik, noKontak, fotoKtp, id)
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
//
//func DeleteUser(id int) (Response, error) {
//	var res Response
//
//	con := db.CreateCon()
//
//	sqlStatement := "UPDATE user SET deleteAt=NOW() WHERE idUser=? AND deleteAt IS NULL"
//	stmt, err := con.Prepare(sqlStatement)
//	if err != nil {
//		return res, err
//	}
//
//	result, err := stmt.Exec(id)
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
