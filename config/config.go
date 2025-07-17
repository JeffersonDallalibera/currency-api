package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

var (
	APIBaseURL string
	APIKey     string
)

func Load() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	APIBaseURL = os.Getenv("API_BASE_URL")
	APIKey = os.Getenv("API_KEY")
}
