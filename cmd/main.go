package main

import (
	"log"
	"net/http"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/database/s3"
	"mybadges/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config loaded: ", cfg)
	s3cfg := cfg.GetS3Config()
	s3storage := s3.New(s3cfg.AccessKey, s3cfg.SecretKey, s3cfg.Endpoint, s3cfg.Bucket, s3cfg.Region)
	
	databaseURL := cfg.GetDatabaseURL()
	storage, err := postgres.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(storage, cfg)

	log.Fatal(http.ListenAndServe(":8080", r))
}
