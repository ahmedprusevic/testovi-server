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
	r.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
	r.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateUser))).Methods("PUT")
	r.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuth(server.DeleteUser)).Methods("DELETE")

	//Login

	r.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	return r

}
