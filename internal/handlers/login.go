package handlers

import (
	"encoding/json"
	stderrors "errors"
	"log"
	"net/http"

	"mybadges/internal/config"
	"mybadges/internal/database"
	"mybadges/internal/database/models"
	"mybadges/internal/utils"
	"mybadges/internal/utils/errors"
)

func Login(userRepo database.UserRepository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println(err)
			http.Error(w, "Incorrect data", http.StatusBadRequest)
			return
		} else {

		}

		err := userRepo.CheckCredentials(user.Email, user.Password)
		if err != nil {
			if stderrors.Is(err, errors.ErrCheckingPassword) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		} else {
			token, err := utils.NewToken(user.Email, cfg.JWT.Lifetime, cfg.JWT.Secret)
			if err != nil {
				log.Println(err)
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			} else {
				w.Header().Set("Authorization", "Bearer "+token)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Authorization successful"))
			}
		}

	}
}
