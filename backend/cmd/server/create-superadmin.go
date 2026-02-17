package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/database"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/utils"
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

	// Create super admin user
	email := "admin@ugenius.com"
	password := "admin123" // Change this in production!

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	superAdmin := &models.User{
		Email:         email,
		PasswordHash:  hashedPassword,
		FirstName:     "Super",
		LastName:      "Admin",
		Role:          models.RoleSuperAdmin,
		EmailVerified: true,
	}

	if err := db.Create(&superAdmin).Error; err != nil {
		log.Fatalf("Failed to create super admin: %v", err)
	}

	fmt.Printf("✅ Super admin created successfully!\n")
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("\n⚠️  Please change the password after first login!\n")
	fmt.Printf("\n🌐 Login at: http://localhost:8080/admin/login\n")
}
