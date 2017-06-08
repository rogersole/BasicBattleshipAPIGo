package router

import (
	"net/http"
	"github.com/rogersole/simple_api/cmd/simple_api/handler"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"Initialize",
		"POST",
		"/game",
		handler.InitializeGame,
	},
	Route{
		"GameTurn",
		"PUT",
		"/game",
		handler.UpdateTurn,
	},
}
