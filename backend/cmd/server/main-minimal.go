package main

import (
	"log"
	"os"
	"path/filepath"

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
	envLoaded := false

	// Try current directory first
	if err := godotenv.Load(); err == nil {
		log.Println("Loaded .env from current directory")
		envLoaded = true
	}

	// Try backend/.env (if running from project root)
	if !envLoaded {
		if err := godotenv.Load("backend/.env"); err == nil {
			log.Println("Loaded .env from backend/.env")
			envLoaded = true
		}
	}

	// Try other paths
	if !envLoaded {
		if exe, err := os.Executable(); err == nil {
			dir := filepath.Dir(exe)
			// Try exe directory
			if err := godotenv.Load(filepath.Join(dir, ".env")); err == nil {
				log.Println("Loaded .env from executable directory")
				envLoaded = true
			}
			// Try going up to backend directory
			if !envLoaded {
				if err := godotenv.Load(filepath.Join(dir, "../../.env")); err == nil {
					log.Println("Loaded .env from parent directory")
					envLoaded = true
				}
			}
		}
	}

	if !envLoaded {
		log.Println("No .env file found, using environment variables")
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
	galleryHandler := handlers.NewGalleryHandler(db)

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

		// Protected routes (require authentication)
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			protected.GET("/me", authHandler.GetMe)
		}

		// Admin-only routes (require authentication + admin role)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		admin.Use(middleware.RequireAdmin())
		{
			admin.GET("/users", authHandler.GetUsers)
			admin.DELETE("/users/:id", authHandler.DeleteUser)
			admin.POST("/gallery/upload", galleryHandler.UploadGalleryImage)
			admin.POST("/gallery", galleryHandler.CreateGallery)
			admin.DELETE("/gallery/:id", galleryHandler.DeleteGallery)
		}

		// Public campus list
		api.GET("/campuses", campusHandler.List)
		api.GET("/campuses/:id", campusHandler.GetByID)

		// Public gallery
		api.GET("/gallery", galleryHandler.GetGallery)

		// Serve static files (uploads) - specific route for gallery images
		router.GET("/uploads/gallery/:filename", func(c *gin.Context) {
			filename := c.Param("filename")
			filePath := "C:/Users/USER/Desktop/u-genius/backend/uploads/gallery/" + filename
			c.File(filePath)
		})

		// Contact form (public - no auth required)
		api.POST("/contact", contactHandler.SendContactEmail)
	}

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
