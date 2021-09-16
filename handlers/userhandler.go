package handlers

import (
	"encoding/json"
	"net/http"

	"../migrations"
	"../models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	migrations.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
