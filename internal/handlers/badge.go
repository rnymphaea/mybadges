package handlers

import (
	"log"
	"net/http"

	"mybadges/internal/config"
	"mybadges/internal/database"
	"mybadges/internal/database/models"
	"mybadges/internal/utils"
	spec "mybadges/internal/utils/badge"
)

func UploadBadge(badgeRepo database.BadgeRepository, imageRepo database.ImageRepository, usersRepo database.UserRepository, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Получение файла
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Unable to get file from form", http.StatusBadRequest)
			return
		}
		defer file.Close()

		badge := models.Badge{
			ID:          database.GenerateUUID(),
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			//CategoryID:   uuid.MustParse(r.FormValue("category_id")),
			//CollectionID: uuid.MustParse(r.FormValue("collection_id")),
			Material: spec.GetMaterial(r.FormValue("material")),
			//Weight:       parseInt(r.FormValue("weight")),
			//Height:       parseInt(r.FormValue("height")),
			//Width:        parseInt(r.FormValue("width")),
			//Thickness:    parseInt(r.FormValue("thickness")),
			Coverage:  r.FormValue("coverage"),
			Fastening: r.FormValue("fastening"),
		}

		tokenString, err := utils.GetTokenFromHeader(r)
		if err != nil {
			http.Error(w, "Unable to get token from header", http.StatusBadRequest)
			log.Println(err)
			return
		}

		secret := cfg.GetSecretKey()
		email, err := utils.GetEmailFromToken(tokenString, secret)
		if err != nil {
			http.Error(w, "Unable to get email from token", http.StatusBadRequest)
			log.Println(err)
			return
		}

		id, err := usersRepo.GetUserIDByEmail(email)
		if err != nil {
			http.Error(w, "Unable to get user id by email", http.StatusBadRequest)
			log.Println(err)
			return
		}

		badge.UserID = id

		if err = badgeRepo.AddBadge(badge); err != nil {
			http.Error(w, "Unable to add badge", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		url, err := imageRepo.UploadFile(file, badge.Title)
		if err != nil {
			http.Error(w, "Unable to add badge photo", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		badge.ImageURL = url

		log.Println(badge)
	}
}
