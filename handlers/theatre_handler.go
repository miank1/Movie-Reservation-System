package handlers

import (
	"net/http"

	"movie-reservation/models"

	"github.com/gin-gonic/gin"
)

var theaters []models.Theater
var nextTheaterID = 1

func GetTheaters(c *gin.Context) {
	c.JSON(http.StatusOK, theaters)
}

func CreateTheater(c *gin.Context) {
	var theater models.Theater

	if err := c.BindJSON(&theater); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	theater.ID = nextTheaterID
	nextTheaterID++

	theaters = append(theaters, theater)

	c.JSON(http.StatusCreated, theater)
}
