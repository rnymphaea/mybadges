package router

import (
	"github.com/gorilla/mux"

	"mybadges/internal/config"
	"mybadges/internal/database/postgres"
	"mybadges/internal/database/s3"
	"mybadges/internal/handlers"
	"mybadges/internal/middleware"
)

func NewRouter(db *postgres.Storage, cfg *config.Config, s3storage *s3.Storage, mw middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	router.Use(mw.LoggingMW)
	router.HandleFunc("/register", handlers.Register(db)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(db, cfg)).Methods("POST")
	key := cfg.GetSecretKey()
	router.HandleFunc("/addbadge", mw.AuthMW(key)(handlers.UploadBadge(db, s3storage, db, cfg))).Methods("POST")
	return router
}
