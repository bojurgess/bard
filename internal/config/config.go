package config

import (
	"errors"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var AppConfig struct {
	Port                string
	Host                string
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyRedirectUri  string
	DatabaseUrl         string
}

func Load() {
	var err error
	AppConfig.Port, err = getEnv("PORT", "3000")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.Host, err = getEnv("HOST", "localhost")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.SpotifyClientId, err = getEnv("SPOTIFY_CLIENT_ID", "")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.SpotifyClientSecret, err = getEnv("SPOTIFY_CLIENT_SECRET", "")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.SpotifyRedirectUri, err = getEnv("SPOTIFY_REDIRECT_URI", "")
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.DatabaseUrl, err = getEnv("DATABASE_URL", "")
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, defaultValue string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		if defaultValue == "" {
			return "", errors.New("Missing required environment variable: " + key)
		}
		return defaultValue, nil
	}
	return value, nil
}
