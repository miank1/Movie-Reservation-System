package main

import (
	"log"
	"movie-reservation/handlers"
	"movie-reservation/models"
	"movie-reservation/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Init DB
	storage.InitDB()

	// Ensure migrations
	if err := storage.DB.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "Movie Reservation API — healthy")
	})

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/login", handlers.Login)
	}

	// Example protected route (any logged-in user)
	me := r.Group("/me")
	me.Use(handlers.AuthMiddleware())
	{
		me.GET("", func(c *gin.Context) {
			uid := c.GetUint("user_id")
			role := c.GetString("role")
			c.JSON(200, gin.H{
				"user_id": uid,
				"role":    role,
			})
		})
	}

	// Example admin-only route
	admin := r.Group("/admin")
	admin.Use(handlers.AuthMiddleware(), handlers.AdminOnly())
	{
		admin.GET("/whoami", func(c *gin.Context) {
			c.JSON(200, gin.H{"ok": "you are admin"})
		})
	}

	// Public movie routes (browse)
	r.GET("/movies", handlers.ListMovies)
	r.GET("/movies/:id", handlers.GetMovie)

	// Admin-only movie management
	adminMovies := r.Group("/movies")
	adminMovies.Use(handlers.AuthMiddleware(), handlers.AdminOnly())
	{
		adminMovies.POST("", handlers.CreateMovie)
		adminMovies.PUT("/:id", handlers.UpdateMovie)
		adminMovies.DELETE("/:id", handlers.DeleteMovie)
	}

	addr := ":8080"
	log.Printf("Starting server on %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
