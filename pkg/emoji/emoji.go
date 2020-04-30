package emoji

import "twb/pkg/forecast"

var (
	CountriesFETA = map[string]string{
		"🇷🇺": "ru",
		"🇬🇧": "en",
	}

	CountriesFATE = map[string]string{
		"ru": "🇷🇺",
		"en": "🇬🇧",
	}

	Help            = "🆘"
	Globe           = "🌏"
	Info            = "ℹ️"
	Sunrise         = "🌄"
	Sunset          = "🌆"
	Gear            = "⚙️"
	TriangularRuler = "📐"
	Back            = "↩️"

	Icons = map[forecast.ConditionType]string{
		forecast.Thunderstorm: "⛈",
		forecast.Drizzle:      "🌧",
		forecast.Rain:         "💦",
		forecast.Snow:         "🌨",
		forecast.Mist:         "🌫",
		forecast.Smoke:        "🌫",
		forecast.Haze:         "🌫",
		forecast.Dust:         "🌫",
		forecast.Fog:          "🌫",
		forecast.Sand:         "🌫",
		forecast.Ash:          "🌫",
		forecast.Squall:       "💨",
		forecast.Tornado:      "🌪",
		forecast.Clear:        "☀️",
		forecast.Clouds:       "☁️",
	}

	Directions = [16]string{
		"⬆️",
		"↗️",
		"↗️",
		"↗️",
		"➡️",
		"↘️",
		"↘️",
		"↘️",
		"⬇️",
		"↙️",
		"↙️",
		"↙️",
		"⬅️",
		"↖️",
		"↖️",
		"↖️",
	}
)
