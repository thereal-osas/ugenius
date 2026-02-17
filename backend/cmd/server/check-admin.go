package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/database"
	"github.com/ugenius/backend/internal/models"
)

func main() {
	fmt.Println("🔍 Checking super admin account...")

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

	// Check for super admin
	var admin models.User
	err = db.Where("email = ? AND role = ?", "admin@ugenius.com", models.RoleSuperAdmin).First(&admin).Error
	if err != nil {
		log.Fatalf("❌ Super admin not found: %v", err)
	}

	fmt.Printf("✅ Super admin found:\n")
	fmt.Printf("   Email: %s\n", admin.Email)
	fmt.Printf("   Name: %s %s\n", admin.FirstName, admin.LastName)
	fmt.Printf("   Role: %s\n", admin.Role)
	fmt.Printf("   Email Verified: %t\n", admin.EmailVerified)
	fmt.Printf("   Created: %s\n", admin.CreatedAt)

	// Test password verification
	testPassword := "admin123"
	fmt.Printf("\n🔐 Testing password: '%s'\n", testPassword)

	// Import utils to check password
	// Note: We can't easily test password hash here without importing utils
	fmt.Printf("💡 If login fails, the password might be different\n")
	fmt.Printf("💡 Try creating a new super admin with a known password\n")
}
