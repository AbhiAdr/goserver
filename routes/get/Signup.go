package get

import (
	"goserver/handlers"
)

func Signup_page() Route {

	return Route{
		"Signup_page",
		"GET",
		"/signup",
		handlers.Signup_page,
	}
}
