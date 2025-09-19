package handlers

import (
	"movie-reservation/models"
	"movie-reservation/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListMovies(c *gin.Context) {
	var movies []models.Movie
	storage.DB.Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func CreateMovie(c *gin.Context) {
	var newMovie models.Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.DB.Create(&newMovie)
	c.JSON(http.StatusCreated, newMovie)
}

func GetMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var movie models.Movie
	result := storage.DB.First(&movie, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var movie models.Movie
	if err := storage.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	var updated models.Movie
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated.ID = movie.ID
	storage.DB.Save(&updated)
	c.JSON(http.StatusOK, updated)
}

func DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := storage.DB.Delete(&models.Movie{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "movie not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
