package main

import (
	"log"

	"github.com/codacube/go-microservice-banking/app"
	"github.com/codacube/go-microservice-banking/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	logger.Info("Starting the application...")
	app.Start()
}
