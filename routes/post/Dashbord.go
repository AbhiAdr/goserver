package post

import (
	"goserver/handlers"
)

func Dashbord() Route {

	return Route{
		"Dashbord",
		"POST",
		"/Dashbord",
		handlers.Dashbord,
	}
}

func Users() Route {

	return Route{
		"Users",
		"POST",
		"/Users",
		handlers.Users,
	}
}
