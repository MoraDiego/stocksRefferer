package handlers

import (
	"encoding/json"
	"net/http"
	"stocksRefferer/db"
	"stocksRefferer/models"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func GetAccionesHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
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
