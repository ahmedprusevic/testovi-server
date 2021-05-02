package main

import (
	"log"
	"os"

	"github.com/ahmedprusevic/testovi-server/rest"
	"github.com/joho/godotenv"
)

var server = &rest.Server{}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	routes := CreateRoutes()

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Start(":8080", routes)
}
