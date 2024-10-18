package main

import (
	"log"
	"net/http"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	databaseURL := config.GetDatabaseURL(cfg)
	storage, err := postgres.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(storage, cfg)

	log.Fatal(http.ListenAndServe(":8080", r))
}
