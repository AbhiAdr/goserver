package handlers

import (
	"API/models"
	"encoding/json"
	"net/http"
)

func Showcadmin(w http.ResponseWriter, r *http.Request) {
	response := models.Showcadmin()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func Addcadmin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	url := r.FormValue("curl")
	user := r.FormValue("cuname")
	pwd := r.FormValue("cpwd")

	if cookie, err := r.Cookie("id"); err == nil {

		value := make(map[string]int64)

		if err = s.Decode("id", cookie.Value, &value); err == nil {
			response := models.Addcadmin(url, user, pwd, value["id"])
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				panic(err)
			}
		}
	}
}
