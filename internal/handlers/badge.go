package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"mybadges/internal/database"
	"mybadges/internal/database/models"
	spec "mybadges/internal/utils/badge"
)

func UploadBadge(badgeRepo database.BadgeRepository, imageRepo database.ImageRepository) http.HandlerFunc {
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
			ID: uuid.New(), // Генерация нового ID
			//UserID:      uuid.MustParse(r.FormValue("user_id")),
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
			//ReleaseDate: r.FormValue("release_date"),
			//ImageURL:     imageURL,
			//Price:        r.FormValue("price"),
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

		if err = badgeRepo.AddBadge(badge); err != nil {
			http.Error(w, "Unable to add badge", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		_, err = imageRepo.UploadFile(file, badge.Title)
		if err != nil {
			http.Error(w, "Unable to add badge photo", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(badge)
	}
}
