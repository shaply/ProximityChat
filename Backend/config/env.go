package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDB_URI            string
	PublicHost             string
	Port                   string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

// Holds the environment variables
var Envs = initConfig()

func initConfig() Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return Config{
		MongoDB_URI:            getEnv("MONGODB_URI", ""),
		PublicHost:             getEnv("PUBLIC_HOST", "localhost"),
		Port:                   getEnv("PORT", "8080"),
		JWTSecret:              getEnv("JWT_SECRET", "not_a_secret"),
		JWTExpirationInSeconds: getEnvasInt64("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvasInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
