package handlers

import (
	"github.com/gorilla/securecookie"
	"html/template"
)

var tpl = template.Must(template.ParseFiles("views/index.html"))

var s *securecookie.SecureCookie

func SecureCookies() {

	s = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

}
