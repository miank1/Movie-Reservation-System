package handlers

import (
	"net/http"

	"movie-reservation/models"

	"github.com/gin-gonic/gin"
)

var screens []models.Screen
var nextScreenID = 1

// Get all screens
func GetScreens(c *gin.Context) {
	c.JSON(http.StatusOK, screens)
}

// Create a new screen
func CreateScreen(c *gin.Context) {

	var screen models.Screen

	// bind JSON request
	if err := c.BindJSON(&screen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// validate theater exists
	theaterExists := false

	for _, t := range theaters {
		if t.ID == screen.TheaterID {
			theaterExists = true
			break
		}
	}

	if !theaterExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "theater does not exist",
		})
		return
	}

	// generate screen ID
	screen.ID = nextScreenID
	nextScreenID++

	// save screen
	screens = append(screens, screen)

	c.JSON(http.StatusCreated, screen)
}
