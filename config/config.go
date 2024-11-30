package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config function to get the value of the key from the .env file
func Config(key string) string {
	err := godotenv.Load(".env")

	// If there is an error loading the .env file
	if err != nil {
		panic("Error loading .env file")
	}

	// Return the value of the key
	return os.Getenv(key)
}
