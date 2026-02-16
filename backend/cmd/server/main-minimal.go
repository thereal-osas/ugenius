package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/database"
	"github.com/ugenius/backend/internal/handlers"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close(db)

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Set Gin mode
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize essential services
	emailService := services.NewEmailService(&cfg.SMTP, cfg.Frontend.URL)
	authService := services.NewAuthService(db, cfg, emailService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	contactHandler := handlers.NewContactHandler(emailService)
	campusHandler := handlers.NewCampusHandler(db)

	// Create router
	router := gin.New()
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.CORSMiddleware(cfg.CORS.Origins))

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "version": "1.0.0"})
		})

		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/register-admin", authHandler.RegisterAdmin)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// Public campus list
		api.GET("/campuses", campusHandler.List)
		api.GET("/campuses/:id", campusHandler.GetByID)

		// Contact form (public - no auth required)
		api.POST("/contact", contactHandler.SendContactEmail)
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
