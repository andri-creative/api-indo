package handlers

import (
	"encoding/json"
	"net/http"

	"api-indo-golang/database"
	"api-indo-golang/models"
)

func GetRegencies(w http.ResponseWriter, r *http.Request) {
	provinceID := r.URL.Query().Get("province_id")

	rows, err := database.DB.Query(
		"SELECT id, province_id, name FROM regencies WHERE province_id = ?",
		provinceID,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var data []models.Regency
	for rows.Next() {
		var r models.Regency
		rows.Scan(&r.ID, &r.ProvinceID, &r.Name)
		data = append(data, r)
	}

	json.NewEncoder(w).Encode(data)
}
