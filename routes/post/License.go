package post

import (
	"go_mjolnir/handlers"
)

func CreateDemoLicense() Route {

	return Route{
		"CreateDemoLicense",
		"POST",
		"/License",
		handlers.CreateDemoLicense,
	}
}

func SubmitLicense() Route {

	return Route{
		"SubmitLicense",
		"POST",
		"/SubmitLicense",
		handlers.SubmitLicense,
	}
}
