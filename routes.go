package main

import (
	"goserver/routes/get"
	"goserver/routes/post"
)

var getroutes = GRoutes{
	get.Index_page(),
	get.Signup_page(),
}

var postroutes = PRoutes{
	post.Login(),
	post.Signup(),
	post.Dashbord(),
	post.Users(),
	post.Logout(),
}
