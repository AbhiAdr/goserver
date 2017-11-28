package handlers

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"goserver/models"
	"goserver/structs"
	"net/http"
)

func Login_page(w http.ResponseWriter, r *http.Request) {

	//	check if user is logged in or Not
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]int64)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			// passing values to page > Check respective struct in ../structs/index_page.go
			d := structs.IndexData{"Index", value["id"], value["login_status"]}
			tpl.ExecuteTemplate(w, "index.html", d)
		}
	} else {
		// passing the page name while rendering > Check respective struct in ../structs/index_page.go
		d := structs.Data{"Login"}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		tpl.ExecuteTemplate(w, "index.html", d)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	var chk_st int

	r.ParseForm()
	email := r.FormValue("email")
	pwd := r.FormValue("pwd")

	data := models.Login(email)
	pwdfromDb := []byte(data.Data.Pwd)

	if err := bcrypt.CompareHashAndPassword(pwdfromDb, []byte(pwd)); err != nil {

		chk_st = 1
	}

	if chk_st == 1 {

		data := models.Login_err()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {

		value := map[string]int64{
			"id":           data.Login_id,
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

func Dashbord(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]int64)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			data := models.LoginDashbord(value["id"])
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				panic(err)
			}
		}
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]int64)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			data := models.ShowUsers()
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				panic(err)
			}
		}
	}
}
