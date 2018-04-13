package main

import (
	"github.com/gorilla/mux"
)

// Router retireve the router
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Ver},
	Route{"Tasks", "GET", "/tasks", TaskList},
	Route{"Tasks.Add", "POST", "/tasks", TaskAdd},
	Route{"Tasks", "GET", "/tasks/{id}", TaskByID},
}
