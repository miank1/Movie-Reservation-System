package handlers

import (
	"movie-reservation/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var movies = []models.Movie{
	{
		ID:       1,
		Title:    "Interstellar",
		Duration: 169,
	},
	{
		ID:       2,
		Title:    "Inception",
		Duration: 148,
	},
}

var nextMovieID = 3

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func GetMovieByID(c *gin.Context) {
	id := c.Param("id")

	for _, movie := range movies {
		if strconv.Itoa(movie.ID) == id {
			c.JSON(http.StatusOK, movie)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie

	// Binding the json to struct
	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	// generate ID
	movie.ID = nextMovieID
	nextMovieID++

	// add movie to slice
	movies = append(movies, movie)

	c.JSON(http.StatusCreated, movie)
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	for i, movie := range movies {
		if strconv.Itoa(movie.ID) == id {

			// remove movie from slice
			movies = append(movies[:i], movies[i+1:]...) // imp

			c.JSON(http.StatusOK, gin.H{
				"message": "movie deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "movie not found",
	})
}
