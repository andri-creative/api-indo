// handlers/import_simple.go
package handlers

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"api-indo-golang/database"
)

func ImportSimpleCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil parameter level
	level := r.URL.Query().Get("level")
	if level == "" {
		http.Error(w, "Level parameter is required: province, regency, district, village", http.StatusBadRequest)
		return
	}

	// Validasi level
	switch level {
	case "province", "regency", "district", "village":
		// OK
	default:
		http.Error(w, "Invalid level. Must be: province, regency, district, village", http.StatusBadRequest)
		return
	}

	// Parse file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Parse CSV
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.TrimLeadingSpace = true

	// Baca header (skip)
	_, err = reader.Read()
	if err != nil && err != io.EOF {
		http.Error(w, "Failed to read CSV", http.StatusBadRequest)
		return
	}

	// Prepare statement berdasarkan level
	var stmt *sql.Stmt
	switch level {
	case "province":
		stmt, err = database.DB.Prepare("INSERT INTO provinces (id, name) VALUES (?, ?)")
	case "regency":
		stmt, err = database.DB.Prepare("INSERT INTO regencies (id, province_id, name) VALUES (?, ?, ?)")
	case "district":
		stmt, err = database.DB.Prepare("INSERT INTO districts (id, regency_id, name) VALUES (?, ?, ?)")
	case "village":
		stmt, err = database.DB.Prepare("INSERT INTO villages (id, district_id, name) VALUES (?, ?, ?)")
	}

	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Import data
	count := 0
	errors := []string{}

	for lineNum := 2; ; lineNum++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, fmt.Sprintf("Line %d: CSV error", lineNum))
			continue
		}

		// Clean data
		for i := range record {
			record[i] = strings.TrimSpace(record[i])
		}

		// Eksekusi berdasarkan level
		var errExec error

		if level == "province" {
			// Provinsi: butuh 2-3 kolom
			if len(record) >= 2 {
				code := record[0]
				name := record[1]

				// Jika ada 3 kolom, nama ada di kolom 3
				if len(record) >= 3 {
					name = record[2]
				}

				// Validasi untuk provinsi: skip jika parent_code bukan 0
				if len(record) >= 3 {
					parentCode := record[1]
					if parentCode != "0" {
						continue // Skip non-province data
					}
				}

				_, errExec = stmt.Exec(code, name)
			}
		} else {
			// Regency, District, Village: butuh 3 kolom
			if len(record) >= 3 {
				code := record[0]
				parentCode := record[1]
				name := record[2]

				// Validasi: skip jika parent_code = 0 untuk non-province
				if parentCode == "0" {
					continue
				}

				_, errExec = stmt.Exec(code, parentCode, name)
			}
		}

		if errExec != nil {
			// Skip duplicate
			if strings.Contains(errExec.Error(), "UNIQUE") ||
				strings.Contains(errExec.Error(), "Duplicate") ||
				strings.Contains(errExec.Error(), "constraint") {
				continue
			}
			errors = append(errors, fmt.Sprintf("Line %d: %v", lineNum, errExec))
			continue
		}

		count++
	}

	// Response
	response := map[string]interface{}{
		"status":   "success",
		"level":    level,
		"filename": fileHeader.Filename,
		"imported": count,
		"errors":   len(errors),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
