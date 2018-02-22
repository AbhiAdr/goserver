package handlers

import (
	"encoding/json"
	"fmt"
	"goserver/models"
	"net/http"
)

func Dashbord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : Dashbord :  ")
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			data, _ := models.Dashbord(value["id"])
			fmt.Println("Dashbord Page -> Data Featch -> send json encode response")
			fmt.Println(data)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				panic(err)
			}
		}
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : Dashbord: ")
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("id", cookie.Value, &value); err == nil {
			data, _ := models.Users()
			fmt.Println("Users LIMIT 10 -> Data Featch -> send json encode response")
			fmt.Println(data)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				panic(err)
			}
		} else {
			fmt.Println(err)
		}
	}
}
