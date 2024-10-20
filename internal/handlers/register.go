package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"mybadges/internal/database"
	"mybadges/internal/database/models"
)

func Register(userRepo database.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println(err)
			http.Error(w, "Incorrect data", http.StatusBadRequest)
			return
		} else {
			user.ID = database.GenerateUUID()
			user.CreatedAt = time.Now()
		}

		if err := userRepo.CreateUser(user); err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}
