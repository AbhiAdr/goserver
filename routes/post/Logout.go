package post

import (
	"goserver/handlers"
)

func Logout() Route {

	return Route{
		"Logout",
		"POST",
		"/Logout",
		handlers.Logout,
	}
}
