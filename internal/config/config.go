package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DataBaseURL string
	Port        string
	JWTSecret   string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	dataBaseURL := os.Getenv("DATABASE_URL")

	port := os.Getenv("PORT")
	JWTSecret := os.Getenv("JWT_SECRET")

	return Config{
		DataBaseURL: dataBaseURL,
		Port:        port,
		JWTSecret:   JWTSecret,
	}

}
