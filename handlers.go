package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	message := r.FormValue("message")

	var recipients []Recipient
	err := db.Select(&recipients, "SELECT phone, email FROM recipients")
	if err != nil {
		log.Printf("Error getting recipients from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	for _, recipient := range recipients {
		go sendMessage(recipient, message)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
