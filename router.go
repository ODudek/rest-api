package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router{
	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Todos",
		"GET",
		"/todos",
		GetTodos,
	},
	Route{
		"Todo",
		"GET",
		"/todos/{id}",
		GetTodo,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todos/{id}",
		CreateTodo,
	},
	Route{
		"DeleteTodo",
		"DELETE",
		"/todos/{id}",
		DeleteTodo,
	},
}
