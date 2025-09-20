package storage

import (
	"log"
	"os"

	"movie-reservation/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=root password=secret dbname=moviedb port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate models
	if err := DB.AutoMigrate(&models.User{}, &models.Movie{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// Seed initial admin (if not exists)
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPass := os.Getenv("ADMIN_PASSWORD")

	// if adminEmail == "" {
	// 	adminEmail = "admin@example.com"
	// }
	// if adminPass == "" {
	// 	adminPass = "adminpass"
	// }

	var count int64
	DB.Model(&models.User{}).Where("email = ?", adminEmail).Count(&count)
	if count == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(adminPass), bcrypt.DefaultCost)
		admin := models.User{
			Email:    adminEmail,
			Password: string(hashed),
			Role:     "admin",
		}
		if err := DB.Create(&admin).Error; err != nil {
			log.Fatalf("failed to create seed admin: %v", err)
		}
		log.Printf("Seeded admin user: %s (password from ADMIN_PASSWORD or default)", adminEmail)
	}
}
