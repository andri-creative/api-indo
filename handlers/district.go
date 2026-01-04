package handlers

import (
	"encoding/json"
	"net/http"

	"api-indo-golang/database"
	"api-indo-golang/models"
)

func GetDistricts(w http.ResponseWriter, r *http.Request) {
	regencyID := r.URL.Query().Get("regency_id")

	rows, err := database.DB.Query(
		"SELECT id, regency_id, name FROM districts WHERE regency_id = ?",
		regencyID,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var data []models.District
	for rows.Next() {
		var d models.District
		rows.Scan(&d.ID, &d.RegencyID, &d.Name)
		data = append(data, d)
	}

	json.NewEncoder(w).Encode(data)
}
