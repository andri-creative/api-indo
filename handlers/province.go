package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"api-indo-golang/database"
	"api-indo-golang/models"
)

func GetProvinces(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		"SELECT id, name FROM provinces",
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var data []models.Province
	for rows.Next() {
		var p models.Province
		rows.Scan(&p.ID, &p.Name)
		data = append(data, p)
	}

	json.NewEncoder(w).Encode(data)
}

func GetProvincesByID(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		"SELECT id, name FROM provinces WHERE id = ?",
		strings.TrimPrefix(r.URL.Path, "/provinces/"),
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var data models.Province
	for rows.Next() {
		rows.Scan(&data.ID, &data.Name)
	}

	json.NewEncoder(w).Encode(data)
}
