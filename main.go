package main

import (
	"API/handlers"
	"API/models"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	// Init Db And Secure Cookies
	models.InitDB("root:root@/project?charset=utf8")
	handlers.SecureCookies()
	// Set Static Files Path
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js/")))
	router.PathPrefix("/js/").Handler(js)
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css/")))
	router.PathPrefix("/css/").Handler(css)
	img := http.StripPrefix("/img/", http.FileServer(http.Dir("./public/img/")))
	router.PathPrefix("/img/").Handler(img)
	fonts := http.StripPrefix("/fonts/", http.FileServer(http.Dir("./public/fonts/")))
	router.PathPrefix("/fonts/").Handler(fonts)

	log.Fatal(http.ListenAndServe(":8086", router))
}
