package main

import (
	"auth_service/pkg/logger"
	"log"
)

func main() {
	logger, file, err := logger.New()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer file.Close()

	log.Println("Starting logging")
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Error("Error message")
	log.Println("Logging completed")

	// Ensure all logs are flushed to the file
	file.Sync()
	log.Println("Logs flushed")

}
