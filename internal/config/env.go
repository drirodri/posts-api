package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port string
}

type AppConfig struct {
	Database DatabaseConfig
	Server ServerConfig
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg := &AppConfig{
		Database: DatabaseConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			Username: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			DBName:   os.Getenv("DATABASE_NAME"),
			SSLMode:  os.Getenv("DATABASE_SSLMODE"),
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
	}

	if cfg.Database.Host == "" || cfg.Database.Port == "" || cfg.Database.Username == "" ||
		cfg.Database.Password == "" || cfg.Database.DBName == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}