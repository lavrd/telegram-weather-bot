package forecast

import (
	"github.com/mlbright/forecast/v2"
	c "github.com/spacelavr/telegram-weather-bot/config"
	"github.com/spacelavr/telegram-weather-bot/utils/errors"
	u "github.com/spacelavr/telegram-weather-bot/utils/format"
)

// GetForecast get weather forecast
func GetForecast(lat, lng float64, lang, units string) *forecast.Forecast {
	f, err := forecast.Get(
		c.Cfg.DarkskyToken,
		u.FTS6(lat),
		u.FTS6(lng),
		"now",
		forecast.Units(units),
		forecast.Lang(lang),
	)
	errors.Check(err)

	return f
}
