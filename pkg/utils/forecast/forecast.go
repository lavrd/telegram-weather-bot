package forecast

import (
	"github.com/mlbright/forecast/v2"
	"github.com/spacelavr/telegram-weather-bot/pkg/config"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/errors"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/format"
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
