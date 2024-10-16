package router

import (
	"github.com/gorilla/mux"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/handlers"
)

func NewRouter(db *postgres.Storage, cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db, cfg)).Methods("POST")
	return router
}
