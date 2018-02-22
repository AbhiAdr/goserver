package models

import (
	"errors"
	"fmt"
	"go_mjolnir/structs"
)

func Login_err() structs.Login_err {
	var data structs.Login_err
	data = structs.Login_err{"F", "Invalid Email / Password"}
	return data
}

func User_get_by_username(data structs.User) (structs.Login, error) {
	var user structs.Login
	stmt, err := db.Prepare("SELECT id, firstname, lastname, email, password, login_status FROM tbl_users where email= ?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(data.Email)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&data.Id, &data.Firstname, &data.Lastname, &data.Email, &data.Pwd, &data.Login_status)

		checkErr(err)
		UpdateUserLoginStatus(data)
		user = structs.Login{"S", data, data.Id}
		//data = structs.Login{Status: "S", Data: data, Login_id: data.Id}
	}

	if err != nil {
		return user, errors.New("Unable to Login")
	}
	return user, nil

}

/*
func Login(email string) structs.Login {

	var data structs.Login

	stmt, err := db.Prepare("SELECT id, firstname, lastname, email, password, login_status FROM tbl_users where email= ?")
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
*/
func UpdateUserLoginStatus(data structs.User) {

	data.Login_status = "1"

	stmt, err := db.Prepare("UPDATE tbl_users SET login_status = ? WHERE id = ?")
	checkErr(err)

	res, err := stmt.Exec(data.Login_status, data.Id)
	checkErr(err)

	fmt.Println("User Login Status updated : ")
	fmt.Println(res)
}
