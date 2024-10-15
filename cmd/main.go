package main

import (
	"fmt"
	"log"
	"net/http"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/router"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", index)
	cfg, err := config.LoadConfig()
	log.Println(cfg.JWT)
	if err != nil {
		log.Fatal(err)
	}
	databaseURL := config.GetDatabaseURL(cfg)
	storage, err := postgres.New(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	r := router.NewRouter(storage)
	log.Fatal(http.ListenAndServe(":8080", r))
}
