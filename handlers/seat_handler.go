package handlers

import (
	"net/http"

	"movie-reservation/models"

	"github.com/gin-gonic/gin"
)

var seats = []models.Seat{}
var nextSeatID = 1

// Get all seats
func GetSeats(c *gin.Context) {
	c.JSON(http.StatusOK, seats)
}

// Create seat
func CreateSeat(c *gin.Context) {

	var seat models.Seat

	if err := c.BindJSON(&seat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	// check if screen exists
	screenExists := false

	for _, s := range screens {
		if s.ID == seat.ScreenID {
			screenExists = true
			break
		}
	}

	if !screenExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "screen not found",
		})
		return
	}

	seat.ID = nextSeatID
	nextSeatID++

	seats = append(seats, seat)

	c.JSON(http.StatusCreated, seat)
}
