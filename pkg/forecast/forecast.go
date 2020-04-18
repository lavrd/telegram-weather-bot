package forecast

import "telegram-weather-bot/pkg/types"

type Forecast interface {
	GetCurrent() (*types.ForecastData, error)
	GetDay() (*types.ForecastData, error)
	GetTomorrow() (*types.ForecastData, error)
	GetWeek() (*types.ForecastData, error)
}
