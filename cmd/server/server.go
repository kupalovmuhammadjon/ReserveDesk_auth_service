package main

import (
	"auth_service/api"
	"auth_service/config"
	"auth_service/pkg/logger"
	"auth_service/storage/postgres"
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
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	router := api.NewRouter(db)
	router.Run(cfg.HTTP_PORT)
}
