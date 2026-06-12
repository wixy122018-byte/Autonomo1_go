package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env, se utilizarán variables del sistema")
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func GetServerPort() string {
	return GetEnv("PORT", "8080")
}

func GetDatabaseDSN() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		return databaseURL
	}

	return GetEnv("DATABASE_DSN", "host=localhost user=postgres password=qwerty dbname=digital_library port=8080 sslmode=disable")
}
