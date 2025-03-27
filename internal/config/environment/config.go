// Package environment provides configuration for different deployment environments
package environment

import (
	"log"
	"os"
	"strings"

	environment "github.com/melegattip/financial-resume-engine/internal/config/environment/constants"
)

// SetUp initializes the environment configuration based on the current deployment scope
func SetUp() {
	log.Println("Global Environment Setup")

	_ = os.Setenv("APPLICATION", environment.Application)

	scope := os.Getenv("SCOPE")

	goEnvironment := os.Getenv("GO_ENVIRONMENT")
	if goEnvironment == "" {
		goEnvironment = environment.Development
		os.Setenv("GO_ENVIRONMENT", goEnvironment)
	}

	if strings.HasSuffix(goEnvironment, environment.Development) {
		setupDevelopmentEnvironment()
	} else if strings.HasSuffix(scope, environment.Beta) {
		setupBetaEnvironment()
	} else {
		setupProdEnvironment()
	}
}

// setupProdEnvironment configures the environment variables for production deployment
func setupProdEnvironment() {
	apiURL, internalURL := getBaseUrls(os.Getenv("GO_ENVIRONMENT"))

	// URLs
	os.Setenv("API_URL", apiURL)
	os.Setenv("INTERNAL_URL", internalURL)
	os.Setenv("DATABASE_URL", "postgresql://user:password@prod-db.niloft.com:5432/financial_db")
	os.Setenv("REDIS_URL", "redis://prod-redis.niloft.com:6379")

	// Configuraciones adicionales
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("ENABLE_CACHE", "true")
	os.Setenv("MAX_CONNECTIONS", "100")
}

// setupBetaEnvironment configures the environment variables for beta deployment
func setupBetaEnvironment() {
	apiURL, internalURL := getBaseUrls(os.Getenv("GO_ENVIRONMENT"))

	// URLs
	os.Setenv("API_URL", apiURL)
	os.Setenv("INTERNAL_URL", internalURL)
	os.Setenv("DATABASE_URL", "postgresql://user:password@beta-db.niloft.com:5432/financial_db")
	os.Setenv("REDIS_URL", "redis://beta-redis.niloft.com:6379")

	// Configuraciones adicionales
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("ENABLE_CACHE", "true")
	os.Setenv("MAX_CONNECTIONS", "50")
}

// setupDevelopmentEnvironment configures the environment variables for local development
func setupDevelopmentEnvironment() {
	apiURL, internalURL := getBaseUrls(os.Getenv("GO_ENVIRONMENT"))

	// URLs
	os.Setenv("API_URL", apiURL)
	os.Setenv("INTERNAL_URL", internalURL)
	os.Setenv("DATABASE_URL", "postgresql://user:password@localhost:5432/financial_db")
	os.Setenv("REDIS_URL", "redis://localhost:6379")

	// Configuraciones adicionales
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("ENABLE_CACHE", "false")
	os.Setenv("MAX_CONNECTIONS", "20")
}

// getBaseUrls returns the API and internal URLs based on the environment
// Returns two strings: the API URL and the internal URL
func getBaseUrls(goEnvironment string) (string, string) {
	if environment.Production == goEnvironment {
		return "http://internal.niloft.com", "http://internal.niloft.com"
	}

	return "https://internal-api.niloft.com", "https://internal-api.niloft.com"
}
