package handlers

import (
	"encoding/json"
	"net/http"

	"api-indo-golang/database"
	"api-indo-golang/models"
)

func GetVillages(w http.ResponseWriter, r *http.Request) {
	districtID := r.URL.Query().Get("district_id")

	rows, err := database.DB.Query(
		"SELECT id, district_id, name FROM villages WHERE district_id = ?",
		districtID,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var data []models.Village
	for rows.Next() {
		var v models.Village
		rows.Scan(&v.ID, &v.DistrictID, &v.Name)
		data = append(data, v)
	}

	json.NewEncoder(w).Encode(data)
}
