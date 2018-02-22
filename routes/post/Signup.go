package post

import (
	"goserver/handlers"
)

func Signup() Route {

	return Route{
		"Signup",
		"POST",
		"/Signup",
		handlers.Signup,
	}
}
