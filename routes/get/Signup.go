package get

import (
	"go_mjolnir/handlers"
)

func Signup_page() Route {

	return Route{
		"Signup_page",
		"GET",
		"/signup",
		handlers.Signup_page,
	}
}
