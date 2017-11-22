package main

import (
	"API/handlers"
)

var routes = Routes{
	Route{
		"Index_page",
		"GET",
		"/",
		handlers.Login_page,
	},
	Route{
		"Login",
		"POST",
		"/Login",
		handlers.Login,
	},
	Route{
		"Signup_page",
		"GET",
		"/signup",
		handlers.Signup_page,
	},
	Route{
		"Signup",
		"POST",
		"/Signup",
		handlers.Signup,
	},
	Route{
		"Logout",
		"POST",
		"/Logout",
		handlers.Logout,
	},
	Route{
		"Dashbord",
		"POST",
		"/Dashbord",
		handlers.Dashbord,
	},
	Route{
		"Addcadmin",
		"POST",
		"/Addcadmin",
		handlers.Addcadmin,
	},
	Route{
		"Showcadmin",
		"POST",
		"/Showcadmin",
		handlers.Showcadmin,
	},
}
