package handlers

import (
	"encoding/json"
	"net/http"

	"go_mjolnir/models"
	"go_mjolnir/structs"
)

func Signup_page(w http.ResponseWriter, r *http.Request) {

	if cookie, err := r.Cookie("id"); err == nil {

		value := make(map[string]string)

		if err = s.Decode("id", cookie.Value, &value); err == nil {

			d := structs.IndexData{"Index", value["id"], value["login_status"]}
			tpl.ExecuteTemplate(w, "index.html", d)
		}

	} else {
		d := structs.Data{"Signup"}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		tpl.ExecuteTemplate(w, "index.html", d)
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fn := r.FormValue("fn")
	ln := r.FormValue("ln")
	email := r.FormValue("email")
	pwd := r.FormValue("pwd")

	chk := models.Singup_chk(email)

	if chk == 1 {
		response := models.Singup_error()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	} else {
		response := models.Singup(fn, ln, email, pwd)

		value := map[string]int64{
			"id":           response.Login_id,
			"login_status": response.Data.Login_status,
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
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	}
}
