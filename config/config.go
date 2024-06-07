package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironment loads the environment variables from the .env file
func Init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
