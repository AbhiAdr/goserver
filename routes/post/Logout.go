package post

import (
	"go_mjolnir/handlers"
)

func Logout() Route {

	return Route{
		"Logout",
		"POST",
		"/Logout",
		handlers.Logout,
	}
}
