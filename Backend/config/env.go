package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	URI string
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
		URI: getEnv("MONGODB_URI"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}
