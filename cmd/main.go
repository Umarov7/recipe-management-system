package main

import (
	"log"
	"recipe-management/internal/api"
	"recipe-management/internal/config"
	"recipe-management/internal/logger"
	"recipe-management/internal/storage/postgres"
)

func main() {
	cfg := config.Load()
	logger := logger.NewLogger()
	storage, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer storage.Close()

	router := api.NewRouter(storage, logger)
	router.Run(cfg.ROUTER_ADDRESS)
}
