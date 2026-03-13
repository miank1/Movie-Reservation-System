package handlers

import (
	"net/http"

	"movie-reservation/models"

	"github.com/gin-gonic/gin"
)

var shows = []models.Show{}
var nextShowID = 1

// Get all shows
func GetShows(c *gin.Context) {
	c.JSON(http.StatusOK, shows)
}

// Create a new show
func CreateShow(c *gin.Context) {

	var show models.Show

	// bind request body
	if err := c.BindJSON(&show); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// check if movie exists
	movieExists := false

	for _, m := range movies {
		if m.ID == show.MovieID {
			movieExists = true
			break
		}
	}

	if !movieExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "movie not found",
		})
		return
	}

	// check if screen exists
	screenExists := false

	for _, s := range screens {
		if s.ID == show.ScreenID {
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

	// generate show id
	show.ID = nextShowID
	nextShowID++

	shows = append(shows, show)

	c.JSON(http.StatusCreated, show)
}
