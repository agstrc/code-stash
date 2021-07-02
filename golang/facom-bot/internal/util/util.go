package util

import (
	"facom-bot/internal/logger"
	"os"
)

// GetEnv returns an environment variable. If the variable is not found, it prints an error message and closes the
// program
func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	logger.Error.Fatalf("Environment variable %s is missing.", key)
	return ""
}
