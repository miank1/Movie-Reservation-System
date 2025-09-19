package main

import (
	"log"
	"movie-reservation/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Movie Reservation API — healthy")
	})

	// movies routes
	r.GET("/movies", handlers.ListMovies)
	r.POST("/movies", handlers.CreateMovie)
	r.GET("/movies/:id", handlers.GetMovie)
	r.PUT("/movies/:id", handlers.UpdateMovie)
	r.DELETE("/movies/:id", handlers.DeleteMovie)

	addr := ":8080"
	log.Printf("Starting server on %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
