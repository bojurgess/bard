package config

import (
	"errors"
	"os"
)

type AppConfig struct {
	Port string
	Host string
}

func Load() (*AppConfig, error) {
	port, err := getEnv("PORT", "3000")
	if err != nil {
		return nil, err
	}
	host, err := getEnv("HOST", "localhost")
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Port: port,
		Host: host,
	}, nil
}

func getEnv(key, defaultValue string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			return "", errors.New("Missing required environment variable: " + key)
		}
		return defaultValue, nil
	}
	return value, nil
}
