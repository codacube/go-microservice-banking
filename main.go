package main

import (
	"github.com/codacube/go-microservice-banking/app"
	"github.com/codacube/go-microservice-banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
