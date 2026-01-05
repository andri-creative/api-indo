package seed

import (
	"api-indo-golang/database"
	"encoding/csv"
	"fmt"
	"os"
)

func ImportRegenciesFromCSV() error {
	file, err := os.Open("data/kabupaten.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Skip header
	for i, row := range records {
		if i == 0 {
			continue
		}

		id := row[0]
		regencyID := row[1]
		name := row[2]

		_, err := database.DB.Exec(
			`INSERT OR IGNORE INTO regencies (id, regency_id, name) VALUES (?, ?, ?)`,
			id,
			regencyID,
			name,
		)
		if err != nil {
			return err
		}
	}

	fmt.Println("✅ Data kabupaten berhasil diimport")
	return nil
}
