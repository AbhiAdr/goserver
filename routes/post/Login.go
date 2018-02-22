package post

import (
	"goserver/handlers"
)

func Login() Route {

	return Route{
		"Login",
		"POST",
		"/Login",
		handlers.Login,
	}
}
