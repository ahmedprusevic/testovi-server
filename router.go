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

	//Groups

	r.HandleFunc("/groups", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.CreateGroup))).Methods("POST")
	r.HandleFunc("/groups", middlewares.SetMiddlewareJSON(server.GetGroups)).Methods("GET")
	r.HandleFunc("/groups/{id}", middlewares.SetMiddlewareJSON(server.GetGroup)).Methods("GET")
	r.HandleFunc("/groups/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateGroup))).Methods("PUT")
	r.HandleFunc("/groups/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.DeleteGroup))).Methods("DELETE")

	//Questions

	r.HandleFunc("/questions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.CreateQuestion))).Methods("POST")
	r.HandleFunc("/questions", middlewares.SetMiddlewareJSON(server.GetQuestions)).Methods("GET")
	r.HandleFunc("/questions/{id}", middlewares.SetMiddlewareJSON(server.GetQuestion)).Methods("GET")
	r.HandleFunc("/questions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateQuestion))).Methods("PUT")
	r.HandleFunc("/questions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.DeleteQuestion))).Methods("DELETE")

	//Tests

	r.HandleFunc("/tests", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.CreateTest))).Methods("POST")
	r.HandleFunc("/tests", middlewares.SetMiddlewareJSON(server.GetTests)).Methods("GET")
	r.HandleFunc("/tests/{id}", middlewares.SetMiddlewareJSON(server.GetTest)).Methods("GET")
	r.HandleFunc("/tests/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.UpdateTest))).Methods("PUT")
	r.HandleFunc("/tests/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(server.DeleteTest))).Methods("DELETE")

	//Login

	r.HandleFunc("/auth/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	return r

}
