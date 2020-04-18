package config

import "os"

type Config struct {
	Google   *Google
	OWM      *OWM
	Telegram *Telegram
	DSN      string
	LogLevel string
}

type Google struct {
	Geocoding *GoogleGeocoding
}

type GoogleGeocoding struct {
	Token string
}

// OpenWeatherMap config
type OWM struct {
	Token string
}

type Telegram struct {
	Token string
	Error *TelegramError
}

type TelegramError struct {
	AdminID string
}

func Parse() *Config {
	return &Config{
		Google: &Google{
			Geocoding: &GoogleGeocoding{
				Token: os.Getenv("GOOGLE_GEOCODING_TOKEN"),
			},
		},
		OWM: &OWM{
			Token: os.Getenv("OWM_TOKEN"),
		},
		Telegram: &Telegram{
			Token: os.Getenv("TELEGRAM_TOKEN"),
			Error: &TelegramError{
				AdminID: os.Getenv("TELEGRAM_ERROR_ADMIN"),
			},
		},
		DSN:      os.Getenv("DSN"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
