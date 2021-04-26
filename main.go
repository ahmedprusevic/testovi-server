package main

import (
	"log"
	"net/http"
)

func main() {
	r := CreateRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
