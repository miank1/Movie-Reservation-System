package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Movie model
type Movie struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
}

var db *gorm.DB

func main() {
	// read DATABASE_DSN from environment
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN not set")
	}

	// connect to postgres
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// migrate schema
	if err := db.AutoMigrate(&Movie{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Postgres connected ✅")
	})

	http.HandleFunc("/movies", handleMovies)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleMovies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var movies []Movie
		if err := db.Find(&movies).Error; err != nil {
			http.Error(w, "failed to fetch movies", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)

	case "POST":
		var m Movie
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		if m.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
			return
		}
		if err := db.Create(&m).Error; err != nil {
			http.Error(w, "failed to create movie", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(m)

	case http.MethodDelete:
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract ID from URL → /movies/{id}
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "invalid URL", http.StatusBadRequest)
			return
		}
		idStr := parts[2]
		id, err := strconv.Atoi(idStr)
		fmt.Println("Id ------>", id)
		if err != nil {
			http.Error(w, "invalid ID", http.StatusBadRequest)
			return
		}

		// Delete movie
		if err := db.Delete(&Movie{}, id).Error; err != nil {
			http.Error(w, "failed to delete movie", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent) // success, no response body

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
