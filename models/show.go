package models

type Show struct {
	ID       int    `json:"id"`
	MovieID  int    `json:"movie_id"`
	ScreenID int    `json:"screen_id"`
	StartAt  string `json:"start_at"`
}
