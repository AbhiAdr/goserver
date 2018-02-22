package handlers

import (
	"encoding/json"
	"fmt"
	"go_mjolnir/models"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : Logout: ")
	if cookie, err := r.Cookie("id"); err == nil {
		fmt.Println(cookie.Value)
		value := make(map[string]string)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			data := models.Logout(value["id"])
			fmt.Println("update the login_status 0 -> send response")
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
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println("cookie not found")
	}
}
