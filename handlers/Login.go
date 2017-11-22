package handlers

import (
	"API/models"
	"API/structs"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login_page(w http.ResponseWriter, r *http.Request) {

	if cookie, err := r.Cookie("id"); err == nil {

		value := make(map[string]int64)

		if err = s.Decode("id", cookie.Value, &value); err == nil {
			d := structs.IndexData{"Index", value["id"], value["jadmin"]}
			tpl.ExecuteTemplate(w, "index.html", d)
		}

	} else {
		d := structs.Data{"Login"}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
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
			"id":     data.Login_id,
			"jadmin": data.Data.Jadmin,
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
