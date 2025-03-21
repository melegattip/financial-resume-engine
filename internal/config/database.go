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

func InitDB() *gorm.DB {
	// Obtener variables de entorno
	dbHost := getEnvOrDefault("DB_HOST", "localhost")
	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("DB_NAME", "financial_resume")
	dbPort := getEnvOrDefault("DB_PORT", "5432")

	// Construir DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Intentar conectar a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Primero, agregar la columna user_id si no existe
	err = db.Exec("DO $$ BEGIN ALTER TABLE transactions ADD COLUMN user_id text; EXCEPTION WHEN duplicate_column THEN NULL; END $$;").Error
	if err != nil {
		log.Printf("Warning: Could not add user_id column: %v", err)
	}

	// Establecer un valor por defecto para registros existentes
	err = db.Exec("UPDATE transactions SET user_id = 'system' WHERE user_id IS NULL").Error
	if err != nil {
		log.Printf("Warning: Could not update existing records: %v", err)
	}

	// Hacer la columna NOT NULL
	err = db.Exec("ALTER TABLE transactions ALTER COLUMN user_id SET NOT NULL").Error
	if err != nil {
		log.Printf("Warning: Could not make user_id NOT NULL: %v", err)
	}

	// Auto-migrar los modelos
	err = db.AutoMigrate(
		&transactions.TransactionModel{},
		&categories.CategoryModel{},
	)
	if err != nil {
		log.Fatal("Error al migrar la base de datos:", err)
	}

	fmt.Println("Database connected successfully")
	return db
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
