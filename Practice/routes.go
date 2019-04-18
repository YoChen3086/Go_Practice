package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"QueryUser",
		"GET",
		"/user",
		QueryUser,
	},
	Route{
		"GetUser",
		"GET",
		"/user/{id}",
		GetUser,
	},
	Route{
		"CreateUser",
		"POST",
		"/user",
		CreateUser,
	},
	Route{
		"UpdateUser",
		"PATCH",
		"/user/{id}",
		UpdateUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{id}",
		DeleteUser,
	},
}
