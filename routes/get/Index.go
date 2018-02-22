package get

import (
	"go_mjolnir/handlers"
)

func Index_page() Route {

	return Route{
		"Index_page",
		"GET",
		"/",
		handlers.Login_page,
	}
}
