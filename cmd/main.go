package main

import (
	"log"
	"net/http"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/database/s3"
	"mybadges/internal/middleware"
	"mybadges/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config loaded")

	s3cfg := cfg.GetS3Config()
	s3storage := s3.New(s3cfg.AccessKey, s3cfg.SecretKey, s3cfg.Endpoint, s3cfg.Bucket, s3cfg.Region)

	databaseURL := cfg.GetDatabaseURL()
	storage, err := postgres.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()
	var mw middleware.Middleware
	mw.LoggingMW = middleware.LoggingMiddleware
	r := router.NewRouter(storage, cfg, s3storage, mw)

	log.Fatal(http.ListenAndServe(":8080", r))
}
