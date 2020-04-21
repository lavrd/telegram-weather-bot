package config

import (
	"errors"
	"os"
)

type Config struct {
	GoogleGeocodingToken string
	OpenWeatherMapToken  string
	TelegramToken        string
	DSN                  string
	LogLevel             string
}

func Parse() (*Config, error) {
	googleGeocodingToken, ok := os.LookupEnv("GOOGLE_GEOCODING_TOKEN")
	if !ok {
		return nil, errors.New("google geocoding token is empty")
	}
	owmToken, ok := os.LookupEnv("OPENWEATHERMAP_TOKEN")
	if !ok {
		return nil, errors.New("open weather map token is empty")
	}
	telegramToken, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		return nil, errors.New("telegram token is empty")
	}
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		return nil, errors.New("data source name is empty")
	}

	return &Config{
		GoogleGeocodingToken: googleGeocodingToken,
		OpenWeatherMapToken:  owmToken,
		TelegramToken:        telegramToken,
		DSN:                  dsn,
		LogLevel:             os.Getenv("LOG_LEVEL"),
	}, nil
}
