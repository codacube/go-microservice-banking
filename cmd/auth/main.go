package main

import (
	"log"

	"github.com/codacube/go-microservice-banking/internal/auth/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	app.Start()
}
