package handlers

import (
	"encoding/json"
	"movie-reservation/models"
	"movie-reservation/storage"
	"net/http"
	"strconv"
	"strings"
)

// ListMovies handles GET /movies
func ListMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Movies)
}

// CreateMovie handles POST /movies
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	newMovie.ID = len(storage.Movies) + 1
	storage.Movies = append(storage.Movies, newMovie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

// GetMovie handles GET /movies/{id}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for _, m := range storage.Movies {
		if m.ID == id {
			json.NewEncoder(w).Encode(m)
			return
		}
	}

	http.Error(w, "movie not found", http.StatusNotFound)
}

// UpdateMovie handles PUT /movies/{id}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var updated models.Movie
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	for i, m := range storage.Movies {
		if m.ID == id {
			updated.ID = id
			storage.Movies[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "movie not found", http.StatusNotFound)
}

// DeleteMovie handles DELETE /movies/{id}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	found := false
	newMovies := []models.Movie{}
	for _, m := range storage.Movies {
		if m.ID == id {
			found = true
			continue
		}
		newMovies = append(newMovies, m)
	}

	if !found {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	storage.Movies = newMovies
	w.WriteHeader(http.StatusNoContent)
}
