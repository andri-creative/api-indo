package seed

import (
	"api-indo-golang/database"
	"encoding/csv"

	// "fmt"
	"log"
	"os"
)

func ImportDistrictsFromCSV() error {
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

	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
		INSERT OR IGNORE INTO districts (id, regency_id, name)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	inserted := 0

	for i, row := range records {
		if i == 0 || len(row) < 3 {
			continue
		}

		_, err := stmt.Exec(row[0], row[1], row[2])
		if err != nil {
			tx.Rollback()
			return err
		}

		inserted++
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Printf("✅ Regencies imported: %d rows\n", inserted)
	return nil
}
