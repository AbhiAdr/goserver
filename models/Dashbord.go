package models

import (
	"errors"
	"fmt"
	"goserver/structs"
)

func Dashbord(id string) (structs.User, error) {

	var data structs.User

	stmt, err := db.Prepare("SELECT id, fn, ln, email, password, login_status FROM users where id= ?")
	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(id)
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		//	var id, login_status string
		//	var fn, ln, email, password string

		//	err = rows.Scan(&id, &fn, &ln, &email, &password, &login_status)
		err = rows.Scan(&data.Id, &data.Firstname, &data.Lastname, &data.Email, &data.Pwd, &data.Login_status)
		checkErr(err)

		//data = structs.UserDetails{id, fn, ln, email, password, login_status}
	}

	if err != nil {
		return data, errors.New("Unable to create user")
	}
	return data, nil
}

func Users() (structs.Users, error) {

	var user structs.User
	var data structs.Users

	rows, err := db.Query("SELECT id, fn, ln, email, login_status FROM users LIMIT 10")
	checkErr(err)

	for rows.Next() {
		//var id, login_status int64
		//var fn, ln, email, password string

		err = rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Login_status)
		checkErr(err)

		data = append(data, user)
	}
	fmt.Println(data)
	if err != nil {
		return data, errors.New("No user listed")
	}
	return data, nil
}

/*
func Users() structs.ShowUserDetails {

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
*/
