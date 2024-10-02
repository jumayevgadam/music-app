package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// LoadConfig is
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("config.LoadConfig.Load: %w", err)
	}

	var c Config
	// Load environment variables into the config struct.
	err := envconfig.Process("my_app", &c)
	if err != nil {
		return nil, fmt.Errorf("internal.config.Process: %v", err)
	}

	// Validate the config
	err = validator.New().Struct(c)
	if err != nil {
		return nil, fmt.Errorf("internal.config.validate: %v", err)
	}

	return &c, nil
}
