package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port         string
	DatabaseURL  string
	JwtSecretKey string
}

func LoadConfig() *AppConfig {
	// Load .env file (optional, safe if not found)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading from environment variables")
	}

	return &AppConfig{
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		JwtSecretKey: getEnv("JWT_SECRET_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
