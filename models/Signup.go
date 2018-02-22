package models

import (
	"go_mjolnir/structs"
	"golang.org/x/crypto/bcrypt"
)

func Singup_error() structs.Signup_err {
	var data structs.Signup_err
	data = structs.Signup_err{"F", "All ready exist"}
	return data
}

func Singup_chk(email string) int {

	stmt, err := db.Prepare("SELECT active FROM users WHERE email=?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(email)
	checkErr(err)
	defer rows.Close()
	var data int
	for rows.Next() {
		var id int

		err = rows.Scan(&id)
		checkErr(err)

		data = id
	}
	return data
}

func Singup(fn string, ln string, email string, pwd string) structs.Signup {

	var data structs.Signup

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	checkErr(err)

	password := string(hash[:])

	stmt, err := db.Prepare("INSERT users SET fn=?,ln=?,email=?,password=?,active=?,login_status=?")
	checkErr(err)

	res, err := stmt.Exec(fn, ln, email, password, 1, 1)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	data = structs.Signup{"S", structs.UserDetails{id, fn, ln, email, password, 1}, id}

	return data
}
