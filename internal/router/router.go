package router

import (
	"github.com/gorilla/mux"

	"mybadges/internal/database/postgres"
	"mybadges/internal/handlers"
)

func NewRouter(db *postgres.Storage) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	return router
}
