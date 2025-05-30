package main

import (
	"database/sql"
	"encoding/json"
	"fireflower/Backend/db"
	"fireflower/Backend/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	router.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
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

func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	result, err := DB.Exec("DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
