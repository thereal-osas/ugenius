package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/database"
)

func main() {
	fmt.Println("🔍 Testing database connection...")

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

	// Print database config (without password)
	fmt.Printf("Database Host: %s\n", cfg.Database.Host)
	fmt.Printf("Database Port: %s\n", cfg.Database.Port)
	fmt.Printf("Database User: %s\n", cfg.Database.User)
	fmt.Printf("Database Name: %s\n", cfg.Database.DBName)
	fmt.Printf("Database Password: %s\n", cfg.Database.Password)
	fmt.Printf("Database SSL Mode: %s\n", cfg.Database.SSLMode)

	// Test connection
	fmt.Println("\n🔌 Connecting to database...")
	dsn := cfg.Database.DSN()
	fmt.Printf("DSN: %s\n", dsn)
	db, err := database.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close(db)

	fmt.Println("✅ Database connection successful!")

	// Test if database exists
	var dbName string
	err = db.Raw("SELECT current_database()").Scan(&dbName).Error
	if err != nil {
		log.Fatalf("❌ Failed to get database name: %v", err)
	}

	fmt.Printf("✅ Connected to database: %s\n", dbName)

	// Test if we can create tables
	fmt.Println("\n🏗️  Testing migrations...")
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	fmt.Println("✅ Migrations successful!")

	// List tables
	var tables []string
	err = db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables).Error
	if err != nil {
		log.Printf("⚠️  Could not list tables: %v", err)
	} else {
		fmt.Printf("✅ Found %d tables\n", len(tables))
		for _, table := range tables {
			fmt.Printf("  - %s\n", table)
		}
	}

	fmt.Println("\n🎉 Database test completed successfully!")
}
