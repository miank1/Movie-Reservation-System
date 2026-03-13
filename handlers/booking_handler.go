package handlers

import (
	"net/http"
	"sync"

	"movie-reservation/models"

	"github.com/gin-gonic/gin"
)

var bookings = []models.Booking{}
var nextBookingID = 1

// mutex for booking
var bookingMutex sync.Mutex

func CreateBooking(c *gin.Context) {

	var booking models.Booking

	if err := c.BindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	// lock
	bookingMutex.Lock()
	defer bookingMutex.Unlock()

	// check if seat already booked
	for _, b := range bookings {
		if b.ShowID == booking.ShowID && b.SeatID == booking.SeatID {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "seat already booked",
			})
			return
		}
	}

	booking.ID = nextBookingID
	nextBookingID++

	bookings = append(bookings, booking)

	c.JSON(http.StatusCreated, booking)
}
