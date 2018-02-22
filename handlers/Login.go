package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"goserver/models"
	"goserver/structs"
)

func Login_page(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login -> Page, Checking for (login=true) dashbord ? login based on cookie value")
	//	check if user is logged in or Not
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			// passing values to page > Check respective struct in ../structs/index_page.go
			d := structs.IndexData{"Index", value["id"], value["login_status"]}
			fmt.Println("Login Page -> Cookie Found -> Dashbord")
			fmt.Println(d)
			tpl.ExecuteTemplate(w, "index.html", d)
		}
	} else {
		fmt.Println("Login Page -> Cookie not Found -> licence found -> Login Page")
		d := structs.Data{"Login"} //	Login Page
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		tpl.ExecuteTemplate(w, "index.html", d)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	var chk_st int
	var user structs.User
	r.ParseForm()
	fmt.Println("POST : LOGIN : Form Value")
	fmt.Println(r.Form)
	user.Email = r.FormValue("email")
	pwd := r.FormValue("pwd")
	//data := models.Login(email)
	data, _ := models.User_get_by_username(user)
	pwdfromDb := []byte(data.Data.Pwd)
	fmt.Println("Featch user details based on input :")
	fmt.Println(data)

	if err := bcrypt.CompareHashAndPassword(pwdfromDb, []byte(pwd)); err != nil {
		chk_st = 1
	}

	fmt.Println("CompareHashAndPassword 0 indicates password matches")
	fmt.Println(chk_st)

	if chk_st == 1 {

		data := models.Login_err()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {

		fmt.Println("Successful login create encryped cookies")
		value := map[string]string{
			"id":           data.Data.Id,
			"login_status": data.Data.Login_status,
		}

		if encoded, err := s.Encode("id", value); err == nil {
			cookie := &http.Cookie{
				Name:  "id",
				Value: encoded,
				Path:  "/",
			}

			http.SetCookie(w, cookie)
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	}
}
