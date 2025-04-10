package main

import (
	"database/sql"
	"encoding/json"
	"fireflower/Backend/db"
	"fireflower/Backend/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var DB *sql.DB

func main() {
	var err error
	DB, err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	router := mux.NewRouter()
	router.HandleFunc("/notes", getNotes).Methods("GET")
	router.HandleFunc("/notes", createNote).Methods("POST")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, title, content, created_at FROM notes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, n)
	}
	json.NewEncoder(w).Encode(notes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	var n models.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := DB.QueryRow(
		"INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at",
		n.Title, n.Content,
	).Scan(&n.ID, &n.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(n)
}
