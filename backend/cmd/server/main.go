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
	// Load .env file - try multiple locations
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

	// Try to find .env relative to executable
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

	// Seed initial data
	if err := database.SeedCampuses(db); err != nil {
		log.Printf("Warning: Failed to seed campuses: %v", err)
	}
	if err := database.SeedEvents(db); err != nil {
		log.Printf("Warning: Failed to seed events: %v", err)
	}

	// Set Gin mode
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize services
	emailService := services.NewEmailService(&cfg.SMTP, cfg.Frontend.URL)
	authService := services.NewAuthService(db, cfg, emailService)
	readingHoursService := services.NewReadingHoursService(db, emailService)
	goalsService := services.NewGoalsService(db)
	achievementsService := services.NewAchievementsService(db, emailService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	campusHandler := handlers.NewCampusHandler(db)
	readingHoursHandler := handlers.NewReadingHoursHandler(readingHoursService)
	goalsHandler := handlers.NewGoalsHandler(goalsService)
	achievementsHandler := handlers.NewAchievementsHandler(achievementsService)
	notificationsHandler := handlers.NewNotificationsHandler(db)
	studyGroupsHandler := handlers.NewStudyGroupsHandler(db)
	scholarshipsHandler := handlers.NewScholarshipsHandler(db)
	waitlistHandler := handlers.NewWaitlistHandler(db)
	eventsHandler := handlers.NewEventsHandler(db)

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

		// Public routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/verify-email", authHandler.VerifyEmail)
			auth.POST("/forgot-password", authHandler.ForgotPassword)
			auth.POST("/reset-password", authHandler.ResetPassword)
			auth.POST("/resend-verification", authHandler.ResendVerification)
		}

		// Public campus list
		api.GET("/campuses", campusHandler.List)
		api.GET("/campuses/:id", campusHandler.GetByID)

		// Public leaderboard
		api.GET("/leaderboard", readingHoursHandler.GetLeaderboard)

		// Waitlist registration (public - no auth required)
		api.POST("/waitlist", waitlistHandler.Register)

		// Public events (no auth required)
		api.GET("/events", eventsHandler.List)
		api.GET("/events/featured", eventsHandler.GetFeatured)
		api.GET("/events/:id", eventsHandler.GetByID)

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			// User profile
			protected.GET("/me", authHandler.GetMe)
			protected.PUT("/me", authHandler.UpdateProfile)
			protected.PUT("/me/password", authHandler.ChangePassword)
			protected.POST("/auth/logout", authHandler.Logout)
			protected.POST("/auth/logout-all", authHandler.LogoutAll)

			// Reading hours
			readingHours := protected.Group("/reading-hours")
			{
				readingHours.POST("", readingHoursHandler.Create)
				readingHours.GET("", readingHoursHandler.List)
				readingHours.GET("/stats", readingHoursHandler.GetStats)
				readingHours.GET("/:id", readingHoursHandler.GetByID)
				readingHours.PUT("/:id", readingHoursHandler.Update)
				readingHours.DELETE("/:id", readingHoursHandler.Delete)
			}

			// Goals
			goals := protected.Group("/goals")
			{
				goals.POST("", goalsHandler.Create)
				goals.GET("", goalsHandler.List)
				goals.GET("/:id", goalsHandler.GetByID)
				goals.PUT("/:id", goalsHandler.Update)
				goals.DELETE("/:id", goalsHandler.Delete)
			}

			// Achievements
			achievements := protected.Group("/achievements")
			{
				achievements.GET("", achievementsHandler.List)
				achievements.GET("/badges", achievementsHandler.GetAllBadges)
				achievements.POST("/check", achievementsHandler.CheckAchievements)
			}

			// Notifications
			notifications := protected.Group("/notifications")
			{
				notifications.GET("", notificationsHandler.List)
				notifications.GET("/unread-count", notificationsHandler.GetUnreadCount)
				notifications.POST("/:id/read", notificationsHandler.MarkAsRead)
				notifications.POST("/read-all", notificationsHandler.MarkAllAsRead)
				notifications.DELETE("/:id", notificationsHandler.Delete)
			}

			// Study Groups
			studyGroups := protected.Group("/study-groups")
			{
				studyGroups.POST("", studyGroupsHandler.Create)
				studyGroups.GET("", studyGroupsHandler.List)
				studyGroups.GET("/:id", studyGroupsHandler.GetByID)
				studyGroups.POST("/:id/join", studyGroupsHandler.Join)
				studyGroups.POST("/:id/leave", studyGroupsHandler.Leave)
			}

			// Scholarships
			scholarships := protected.Group("/scholarships")
			{
				scholarships.GET("", scholarshipsHandler.List)
				scholarships.GET("/my-applications", scholarshipsHandler.MyApplications)
				scholarships.GET("/:id", scholarshipsHandler.GetByID)
				scholarships.POST("/:id/apply", scholarshipsHandler.Apply)
			}

			// Admin routes
			admin := protected.Group("/admin")
			admin.Use(middleware.RequireAdmin())
			{
				admin.GET("/reading-hours", readingHoursHandler.AdminList)
				admin.POST("/reading-hours/:id/review", readingHoursHandler.Review)
				admin.GET("/stats", readingHoursHandler.GetCampusStats)
				admin.POST("/campuses", campusHandler.Create)
				admin.PUT("/campuses/:id", campusHandler.Update)
			}
		}
	}

	// Start server
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
