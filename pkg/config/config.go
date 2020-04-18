package config

import "os"

type Config struct {
	GoogleGeocodingToken string
	OpenWeatherMapToken  string
	TelegramToken        string
	DSN                  string
	LogLevel             string
}

func Parse() *Config {
	return &Config{
		GoogleGeocodingToken: os.Getenv("GOOGLE_GEOCODING_TOKEN"),
		OpenWeatherMapToken:  os.Getenv("OPENWEATHERMAP_TOKEN"),
		TelegramToken:        os.Getenv("TELEGRAM_TOKEN"),
		DSN:                  os.Getenv("DSN"),
		LogLevel:             os.Getenv("LOG_LEVEL"),
	}
}
