package seed

import (
	"api-indo-golang/database"
	"encoding/csv"
	"fmt"
	"os"
)

func ImportProvincesFromCSV() error {
	file, err := os.Open("data/provinsi.csv")
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
		name := row[1]

		_, err := database.DB.Exec(
			`INSERT OR IGNORE INTO provinces (id, name) VALUES (?, ?)`,
			id,
			name,
		)
		if err != nil {
			return err
		}
	}

	fmt.Println("✅ Data provinsi berhasil diimport")
	return nil
}
