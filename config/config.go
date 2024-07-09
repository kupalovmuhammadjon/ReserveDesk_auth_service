package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT         string
	AUTH_SERVICE_PORT string
	DB_HOST           string
	DB_PORT           string
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	SIGNING_KEY       string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", 8080))
	config.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", 7777))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "user_management"))
	config.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "dfghjmhgfd"))

	return config
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
