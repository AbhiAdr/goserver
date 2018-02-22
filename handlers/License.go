package handlers

import (
	"encoding/json"
	"fmt"
	"go_mjolnir/models"
	"go_mjolnir/structs"
	// "io/ioutil"
	"net/http"
	// "reflect"
	"time"
)

func CreateDemoLicense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : License -> CreateDemoLicense")
	lic_, err := models.Encrypt([]byte("DemoVersion"))
	checkErr(err)

	date_ := time.Now().UTC().AddDate(1, 0, 0).Format("2006-01-02")
	exp, err := models.Encrypt([]byte(date_))
	checkErr(err)

	demo_data := structs.LicenseJson{string(lic_), string(exp)}

	fmt.Println("Created Demo License")
	fmt.Println(demo_data)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(demo_data); err != nil {
		panic(err)
	}
}

func SubmitLicense(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST : SubmitLicense")
	r.ParseForm()
	decrypt_license, err := models.Decrypt([]byte(r.Form["demo"][0]))
	checkErr(err)
	demo := string(decrypt_license)
	if demo == "DemoVersion" {
		models.CreateLicenseFile()
	}
	// lic_, err := models.Encrypt([]byte("DemoVersion"))
	// checkErr(err)

	// date_ := time.Now().UTC().AddDate(1, 0, 0).Format("2006-01-02")
	// exp, err := models.Encrypt([]byte(date_))
	// checkErr(err)

	// demo_data := structs.LicenseJson{string(lic_), string(exp)}

	// fmt.Println("Created Demo License")
	// fmt.Println(demo_data)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(demo_data); err != nil {
	// 	panic(err)
	// }
}
