package models

type Movie struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title string `json:"title" binding:"required" gorm:"not null"`
	Year  int    `json:"year" binding:"required,gte=1888" gorm:"not null"`
}
