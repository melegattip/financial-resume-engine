// Package config provides configuration utilities for the application
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/melegattip/financial-resume-engine/internal/usecases/categories"
	"github.com/melegattip/financial-resume-engine/internal/usecases/transactions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes and returns a database connection using environment variables
// It also performs necessary migrations and schema updates
func InitDB() *gorm.DB {
	// Get environment variables
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("DB_NAME", "financial_resume")
	dbPort := getEnvOrDefault("DB_PORT", "5432")

	// Build DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Try to connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// First, add user_id column if it doesn't exist
	err = db.Exec("DO $$ BEGIN ALTER TABLE transactions ADD COLUMN user_id text; EXCEPTION WHEN duplicate_column THEN NULL; END $$;").Error
	if err != nil {
		log.Printf("Warning: Could not add user_id column: %v", err)
	}

	// Set a default value for existing records
	err = db.Exec("UPDATE transactions SET user_id = 'system' WHERE user_id IS NULL").Error
	if err != nil {
		log.Printf("Warning: Could not update existing records: %v", err)
	}

	// Make the column NOT NULL
	err = db.Exec("ALTER TABLE transactions ALTER COLUMN user_id SET NOT NULL").Error
	if err != nil {
		log.Printf("Warning: Could not make user_id NOT NULL: %v", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(
		&transactions.TransactionModel{},
		&categories.CategoryModel{},
	)
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	fmt.Println("Database connected successfully")
	return db
}

// getEnvOrDefault returns the value of an environment variable or a default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
