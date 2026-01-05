package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"api-indo-golang/database"
	"api-indo-golang/handlers"
	"api-indo-golang/seed"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env not found, using system env")
	}

	database.Connect()

	err := seed.ImportDistrictsFromCSV()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Buat server dengan timeout settings
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      nil,               // menggunakan http.DefaultServeMux
		ReadTimeout:  30 * time.Second,  // Timeout untuk membaca request
		WriteTimeout: 120 * time.Second, // Timeout untuk menulis response (perbesar untuk import besar)
		IdleTimeout:  120 * time.Second, // Timeout untuk koneksi idle
	}

	// Setup routes
	http.HandleFunc("/provinces", handlers.GetProvinces)
	http.HandleFunc("/regencies", handlers.GetRegencies)
	http.HandleFunc("/districts", handlers.GetDistricts)
	http.HandleFunc("/villages", handlers.GetVillages)

	http.HandleFunc("/import/simple", handlers.ImportSimpleCSV)

	log.Println("🚀 Server running on :" + port)

	// Jalankan server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
