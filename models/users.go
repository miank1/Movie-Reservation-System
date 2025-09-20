package models

import "time"

// User represents an application user
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null" binding:"required,email"`
	Password  string    `json:"-" gorm:"not null"`                 // hashed password, not returing in JSON
	Role      string    `json:"role" gorm:"not null;default:user"` // "admin" or "user"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
