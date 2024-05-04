package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	connStr := GetDBConnectionString()
	var err error
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/send", handleSend)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
