package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("🗄️  Creating database 'ugenius'...")
	
	// Connect to PostgreSQL server (without specifying database)
	dsn := "host=localhost port=5432 user=postgres password=Mirror1#@ sslmode=disable dbname=postgres"
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()
	
	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Failed to ping PostgreSQL: %v", err)
	}
	
	fmt.Println("✅ Connected to PostgreSQL server")
	
	// Check if database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'ugenius')").Scan(&exists)
	if err != nil {
		log.Fatalf("❌ Failed to check if database exists: %v", err)
	}
	
	if exists {
		fmt.Println("✅ Database 'ugenius' already exists")
	} else {
		// Create database
		_, err = db.Exec("CREATE DATABASE ugenius")
		if err != nil {
			log.Fatalf("❌ Failed to create database: %v", err)
		}
		fmt.Println("✅ Database 'ugenius' created successfully")
	}
	
	fmt.Println("\n🎉 Database setup complete!")
}
