package post

import (
	"go_mjolnir/handlers"
)

func Login() Route {

	return Route{
		"Login",
		"POST",
		"/Login",
		handlers.Login,
	}
}
