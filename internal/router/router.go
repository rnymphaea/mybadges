package router

import (
	"github.com/gorilla/mux"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/database/s3"
	"mybadges/internal/handlers"
)

func NewRouter(db *postgres.Storage, cfg *config.Config, s3storage *s3.Storage) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db, cfg)).Methods("POST")
	router.HandleFunc("/addbadge", handlers.UploadBadge(db, s3storage)).Methods("POST")
	return router
}
