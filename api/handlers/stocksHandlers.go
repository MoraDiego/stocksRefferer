package handlers

import (
	"encoding/json"
	"net/http"
	"stocksRefferer/db"
	"stocksRefferer/models"
	"strconv"
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

func GetAccionesByRiskHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	riskStr := r.URL.Query().Get("risk")
	risk, err := strconv.Atoi(riskStr)
	if err != nil || risk < 0 || risk > 99 {
		http.Error(w, "Invalid risk level (0-3)", http.StatusBadRequest)
		return
	}
	DB := db.Connect()
	var acciones []models.Stock
	result := DB.Raw(
		`SELECT *
        FROM stocks
        WHERE
            CASE rating_to
                WHEN 'Strong Buy' THEN 0
                WHEN 'Top Pick' THEN 0
                WHEN 'Buy' THEN 1
                WHEN 'Outperform' THEN 1
                WHEN 'Overweight' THEN 1
                WHEN 'Sector Outperform' THEN 1
                WHEN 'Moderate Buy' THEN 2
                WHEN 'Speculative Buy' THEN 2
                WHEN 'Mkt Outperform' THEN 2
                WHEN 'Market Outperform' THEN 2
                WHEN 'Hold' THEN 3
                WHEN 'Neutral' THEN 3
                WHEN 'Equal Weight' THEN 3
                WHEN 'In-Line' THEN 3
                WHEN 'Market Perform' THEN 3
                WHEN 'Sector Perform' THEN 3
                WHEN 'Peer Perform' THEN 3
                WHEN 'Sector Weight' THEN 3
                WHEN 'Sell' THEN 4
                WHEN 'Strong Sell' THEN 4
                WHEN 'Underperform' THEN 4
                WHEN 'Sector Underperform' THEN 4
                WHEN 'Reduce' THEN 4
                WHEN 'Underweight' THEN 4
                ELSE 99
            END <= ?`, risk).Scan(&acciones)
	if result.Error != nil {
		http.Error(w, "Error al obtener datos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(acciones)
}
