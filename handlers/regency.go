package handlers

import (
	"encoding/json"
	"net/http"

	"api-indo-golang/database"
	"api-indo-golang/models"
)

func GetRegencies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	provinceID := r.URL.Query().Get("province_id")
	if provinceID == "" {
		http.Error(w, "province_id is required", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT
			p.id,
			p.name,
			r.id,
			r.name
		FROM provinces p
		JOIN regencies r ON r.province_id = p.id
		WHERE p.id = ?
	`, provinceID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var response models.RegencyByProvinceResponse
	var regencies []models.Regency
	var provinceSet bool

	for rows.Next() {
		var (
			provinceID   string
			provinceName string
			regencyID    string
			regencyName  string
		)

		if err := rows.Scan(
			&provinceID,
			&provinceName,
			&regencyID,
			&regencyName,
		); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Set province SEKALI
		if !provinceSet {
			response.Province = models.Province{
				ID:   provinceID,
				Name: provinceName,
			}
			provinceSet = true
		}

		regencies = append(regencies, models.Regency{
			ID:         regencyID,
			ProvinceID: provinceID,
			Name:       regencyName,
		})
	}

	// Jika tidak ada data
	if !provinceSet {
		http.Error(w, "province not found", http.StatusNotFound)
		return
	}

	response.Regencies = regencies
	json.NewEncoder(w).Encode(response)
}
