package main

import (
	"github.com/gin-gonic/gin"

	"movie-reservation/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Movie Reservation Service Running",
		})
	})

	router.GET("/movies", handlers.GetMovies)
	router.GET("/movies/:id", handlers.GetMovieByID)
	router.POST("/movies", handlers.CreateMovie)
	router.DELETE("/movies/:id", handlers.DeleteMovie)

	router.GET("/theaters", handlers.GetTheaters)
	router.POST("/theaters", handlers.CreateTheater)
	router.Run(":8080")
}
