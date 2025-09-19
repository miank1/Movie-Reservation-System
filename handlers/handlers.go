package handlers

import (
	"movie-reservation/models"
	"movie-reservation/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListMovies(c *gin.Context) {
	c.JSON(http.StatusOK, storage.Movies)
}

func CreateMovie(c *gin.Context) {
	var newMovie models.Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	newMovie.ID = len(storage.Movies) + 1
	storage.Movies = append(storage.Movies, newMovie)
	c.JSON(http.StatusCreated, newMovie)
}

func GetMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	for _, m := range storage.Movies {
		if m.ID == id {
			c.JSON(http.StatusOK, m)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
}

func UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updated models.Movie
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	for i, m := range storage.Movies {
		if m.ID == id {
			updated.ID = id
			storage.Movies[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
}

func DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	storage.Movies = newMovies
	c.Status(http.StatusNoContent)
}
