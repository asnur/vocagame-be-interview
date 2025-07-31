package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func NewConfig(target any) error {
	filename := os.Getenv("ENV_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", target); err != nil {
			return fmt.Errorf("%s failed to read from env variable", err)
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return fmt.Errorf("%s failed to read from .env file", err)
	}

	if err := envconfig.Process("", target); err != nil {
		return fmt.Errorf("%s failed to read from env variable", err)
	}

	return nil
}
