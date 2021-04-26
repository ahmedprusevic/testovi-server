package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

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
}

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		var h http.Handler
		h = route.HandlerFunc
		h = Logger(h, route.Name)

		r.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(h)
	}
	return r

}
