package models

type Seat struct {
	ID       int    `json:"id"`
	ScreenID int    `json:"screen_id"`
	Row      string `json:"row"`
	Number   int    `json:"number"`
}
