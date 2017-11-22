package handlers

import (
	"API/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	if cookie, err := r.Cookie("id"); err == nil {

		value := make(map[string]int64)

		if err = s.Decode("id", cookie.Value, &value); err == nil {

			data := models.Logout(value["id"])

			cookie := &http.Cookie{
				Name:   "id",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			}

			http.SetCookie(w, cookie)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				panic(err)
			}
		}

	} else {
		fmt.Println("cookie not found")
	}
}
