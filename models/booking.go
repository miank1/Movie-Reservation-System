package models

type Booking struct {
	ID     int `json:"id"`
	ShowID int `json:"show_id"`
	SeatID int `json:"seat_id"`
}
