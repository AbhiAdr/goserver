package models

import (
	"API/structs"
)

func Showcadmin() structs.CadminDetails {

	var data structs.CadminDetails

	rows, err := db.Query("SELECT id, name, url, active FROM admins LIMIT 10")
	checkErr(err)

	for rows.Next() {
		var id, active int64
		var name, url string

		err = rows.Scan(&id, &name, &url, &active)
		checkErr(err)

		data = append(data, structs.CadminDetail{id, name, url, active})
	}

	return data
}

func Addcadmin(url string, uname string, pwd string, uid int64) structs.Addcadmin {
	var data structs.Addcadmin

	stmt, err := db.Prepare("INSERT admins SET name=?,url=?,password=?,created_by=?,active=?")
	checkErr(err)

	res, err := stmt.Exec(uname, url, pwd, uid, 1)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	data = structs.Addcadmin{"S", structs.CadminDetail{id, uname, url, 1}, id}

	return data
}
