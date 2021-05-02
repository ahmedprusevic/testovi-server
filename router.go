package main

import (
	"github.com/ahmedprusevic/testovi-server/middlewares"
	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Index)).Methods("GET")

	//Users

	r.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	r.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")

	return r

}
