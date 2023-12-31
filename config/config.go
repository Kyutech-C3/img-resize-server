package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_DB       string
)

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("failed load env")
	}
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
}
