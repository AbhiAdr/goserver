package handlers

import (
	"encoding/json"
	"fmt"
	"go_mjolnir/models"
	"net/http"
)

func Testplans(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : Testplans: ")
	if cookie, err := r.Cookie("id"); err == nil {
		value := make(map[string]string)

		if err = s.Decode("id", cookie.Value, &value); err == nil {
			fmt.Println(value["id"])
			data, _ := models.Testplans(value["id"])
			fmt.Println("Testplans -> Data Featch -> send json encode response")
			fmt.Println(data)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				panic(err)
			}
		}
	}
}
