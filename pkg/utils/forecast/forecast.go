package forecast

import (
	"telegram-weather-bot/pkg/config"
	"telegram-weather-bot/pkg/utils/errors"
	"telegram-weather-bot/pkg/utils/format"

	forecast "github.com/mlbright/darksky/v2"
)

func GetForecast(lat, lng float64, lang, units string) *forecast.Forecast {
	f, err := forecast.Get(
		config.Viper.Darksky.Token,
		format.FTS6(lat),
		format.FTS6(lng),
		"now",
		forecast.Units(units),
		forecast.Lang(lang),
	)
	errors.Check(err)

	return f
}
