package models

import (
	"fmt"
	"goserver/structs"
)

func Login_err() structs.Login_err {
	var data structs.Login_err
	data = structs.Login_err{"F", "Invalid Email / Password"}
	return data
}

func Login(email string) structs.Login {

	var data structs.Login

	stmt, err := db.Prepare("SELECT id, fn, ln, email, password, login_status FROM users where email= ?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(email)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, login_status int64
		var fn, ln, email, password string

		err = rows.Scan(&id, &fn, &ln, &email, &password, &login_status)
		checkErr(err)

		data = structs.Login{"S", structs.UserDetails{id, fn, ln, email, password, login_status}, id}
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

	stmt, err := db.Prepare("SELECT id, fn, ln, email, password, login_status FROM users where id= ?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(id)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id, login_status int64
		var fn, ln, email, password string

		err = rows.Scan(&id, &fn, &ln, &email, &password, &login_status)
		checkErr(err)

		data = structs.UserDetails{id, fn, ln, email, password, login_status}
	}

	return data
}

func ShowUsers() structs.ShowUserDetails {

	var data structs.ShowUserDetails

	rows, err := db.Query("SELECT id, fn, ln, email, password, login_status FROM users LIMIT 10")
	checkErr(err)

	for rows.Next() {
		var id, login_status int64
		var fn, ln, email, password string

		err = rows.Scan(&id, &fn, &ln, &email, &password, &login_status)
		checkErr(err)

		data = append(data, structs.UserDetails{id, fn, ln, email, password, login_status})
	}

	return data
}
