package config

import "os"

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
	return GetEnv("DATABASE_DSN", "host=localhost user=postgres password=postgres dbname=biblioteca port=5432 sslmode=disable")
}