package main

import (
	"fmt"
	"log"
	"movie-reservation/handlers"
	"net/http"
)

func main() {
	// health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Movie Reservation API — healthy")
	})

	// movies endpoints
	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ListMovies(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateMovie(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetMovie(w, r)
		case http.MethodPut:
			handlers.UpdateMovie(w, r)
		case http.MethodDelete:
			handlers.DeleteMovie(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := ":8080"
	log.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
