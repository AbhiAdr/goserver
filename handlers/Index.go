package handlers

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"html/template"
	"log"
	"os"
)

var licenseExist int
var tpl = template.Must(template.ParseFiles("views/index.html"))
var s *securecookie.SecureCookie

func SecureCookies() {

	s = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
}

func License() {

	filename := "license.txt"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println(err)
		licenseExist = 0
	} else {
		licenseExist = 1 // file exist
	}
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
