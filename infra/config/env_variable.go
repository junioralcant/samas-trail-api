package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	logger := GetLogger("env_variable")
	err := godotenv.Load(".env")

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
	}

	return os.Getenv(key)
}
