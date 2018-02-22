package main

import (
	"go_mjolnir/routes/get"
	"go_mjolnir/routes/post"
)

var getroutes = GRoutes{
	get.Index_page(),
	get.Signup_page(),
}

var postroutes = PRoutes{
	post.CreateDemoLicense(),
	post.SubmitLicense(),
	post.Login(),
	post.Signup(),
	post.Dashbord(),
	post.Users(),
	post.Testplans(),
	post.Logout(),
}
