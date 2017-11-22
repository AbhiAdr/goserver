package models

import (
	"API/structs"
	"fmt"
)

func Login_err() structs.Login_err {
	var data structs.Login_err
	data = structs.Login_err{"F", "Invalid Email / Password"}
	return data
}

func Login(email string) structs.Login {

	var data structs.Login

	stmt, err := db.Prepare("SELECT id, fn, ln, email, password, jadmin FROM users where email= ?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(email)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, jadmin int64
		var fn, ln, email, password string

		err = rows.Scan(&id, &fn, &ln, &email, &password, &jadmin)
		checkErr(err)

		data = structs.Login{"S", structs.UserDetails{id, fn, ln, email, password, jadmin}, id}
		UpdateUserLoginStatus(id)
	}

	return data
}

func UpdateUserLoginStatus(id int64) {

	stmt, err := db.Prepare("UPDATE users SET login_status = ? WHERE id = ?")
	checkErr(err)

	res, err := stmt.Exec(1, id)
	checkErr(err)

	fmt.Println(res)

}

func LoginDashbord(id int64) structs.UserDetails {

	var data structs.UserDetails

	stmt, err := db.Prepare("SELECT id, fn, ln, email, password, jadmin FROM users where id= ?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(id)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, jadmin int64
		var fn, ln, email, password string

		err = rows.Scan(&id, &fn, &ln, &email, &password, &jadmin)
		checkErr(err)

		data = structs.UserDetails{id, fn, ln, email, password, jadmin}
	}

	return data
}
