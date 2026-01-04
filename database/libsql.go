package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func Connect() {
	url := os.Getenv("LIBSQL_URL")
	token := os.Getenv("LIBSQL_AUTH_TOKEN")

	if url == "" || token == "" {
		log.Fatal("❌ LIBSQL_URL atau LIBSQL_AUTH_TOKEN belum diset")
	}

	var err error
	DB, err = sql.Open("libsql", url+"?authToken="+token)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Connected to Turso (libSQL)")
}
