package config

import (
	"log"

	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Error to open file of configuration: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Error to open file of configuration:: %v", err)
	}
	log.Println("Configuration loaded succesfully.")
}
