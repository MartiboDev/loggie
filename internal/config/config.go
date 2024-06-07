package config

import (
	"os"
)

// LoadEnvironment loads the environment variables from the .env file
func Init() {

}

func Get(key string) string {
	return os.Getenv(key)
}
