package router

import (
	"net/http"

	handlers "github.com/sixlib/go/Handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Auth        bool
	HandlerFunc http.HandlerFunc
}
type Routes []Route

//路由配置
var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		false,
		handlers.IndexHandler,
	},
	Route{
		"todo",
		"POST",
		"/todo",
		false,
		handlers.TodoHandler,
	},
	Route{
		"todobykey",
		"POST",
		"/todo/{key}",
		false,
		handlers.TodoOfKeyHandler,
	},
}
