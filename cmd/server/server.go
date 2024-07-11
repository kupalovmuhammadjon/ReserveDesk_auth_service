package main

import (
	"auth_service/api"
	"auth_service/config"
	"auth_service/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	router := api.NewRouter(db)
	router.Run(cfg.HTTP_PORT)
}
