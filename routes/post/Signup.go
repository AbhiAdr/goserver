package post

import (
	"go_mjolnir/handlers"
)

func Signup() Route {

	return Route{
		"Signup",
		"POST",
		"/Signup",
		handlers.Signup,
	}
}
