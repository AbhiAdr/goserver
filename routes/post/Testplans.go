package post

import (
	"go_mjolnir/handlers"
)

func Testplans() Route {

	return Route{
		"Testplans",
		"POST",
		"/Testplans",
		handlers.Testplans,
	}
}
