package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	} `yaml:"server"`

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig() {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Assign environment variables to the AppConfig structure
	AppConfig.Server.Port = os.Getenv("SERVER_PORT")
	AppConfig.Database.Host = os.Getenv("DB_HOST")
	AppConfig.Database.Port = getEnvAsInt("DB_PORT", 5432) // Default value 5432
	AppConfig.Database.User = os.Getenv("DB_USER")
	AppConfig.Database.Password = os.Getenv("DB_PASSWORD")
	AppConfig.Database.Name = os.Getenv("DB_NAME")

	log.Println("Configuration loaded successfully.")
}

// Function to convert an environment variable to an integer with a default value
func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	// Convert the environment variable to an integer
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting %s to int, using default value %d", key, defaultValue)
		return defaultValue
	}
	return result
}
