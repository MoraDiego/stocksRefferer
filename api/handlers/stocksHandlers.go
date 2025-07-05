package handlers

import (
	"encoding/json"
	"net/http"
	"stocksRefferer/db"
	"stocksRefferer/models"
)

func GetAccionesHandler(w http.ResponseWriter, r *http.Request) {

	DB := db.Connect()

	var acciones []models.Stock
	result := DB.Where("rating_to = ?", "Buy").Find(&acciones)

	if result.Error != nil {
		http.Error(w, "Error al obtener datos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acciones)
}
