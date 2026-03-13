package models

type Screen struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TheaterID int    `json:"theater_id"`
}
