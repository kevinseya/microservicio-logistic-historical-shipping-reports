package database

import (
	"fmt"
	"historical-shipping-reports/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// String ConnectDB
func ConnectDB() {
	cfg := config.AppConfig.Database

	// Cadena de conexión para PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	log.Printf("Connection to PostgreSQL: %s\n", dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect with database: %v", err)
	}
	log.Println("Connection successfully to database.")
}
