package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/database"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/utils"
)

func main() {
	fmt.Println("🔧 Creating verified super admin...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
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

	// Delete existing super admin (force delete)
	fmt.Println("🗑️  Removing existing super admin...")
	db.Unscoped().Where("email = ?", "admin@ugenius.com").Delete(&models.User{})

	// Create new verified super admin
	email := "admin@ugenius.com"
	password := "admin123"

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
		EmailVerified: true, // This is the key fix!
	}

	if err := db.Create(&superAdmin).Error; err != nil {
		log.Fatalf("Failed to create super admin: %v", err)
	}

	fmt.Printf("✅ Verified super admin created successfully!\n")
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Email Verified: %t\n", superAdmin.EmailVerified)
	fmt.Printf("\n🎯 Login credentials:\n")
	fmt.Printf("📧 Email: %s\n", email)
	fmt.Printf("🔑 Password: %s\n", password)
	fmt.Printf("\n🌐 Login at: http://localhost:8080/admin/login\n")
	fmt.Printf("🎉 This should work now!\n")
}
