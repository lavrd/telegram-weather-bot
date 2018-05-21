package weather

import (
	"github.com/spacelavr/telegram-weather-bot/pkg/model"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/forecast"
)

func CurrentWeather(lat, lng float64, location string, user *model.DB) string {
	f := forecast.GetForecast(lat, lng, user.Lang, user.Units)

	return getTime(f.Currently.Time, f.Timezone) +
		getCity(location) + getCurrentWeather(user.Lang, user.Units, f)
}

func CurrentWeatherFromLocation(lat, lng float64, location string, user *model.DB) string {
	f := forecast.GetForecast(lat, lng, user.Lang, user.Units)

	return getTime(f.Currently.Time, f.Timezone) +
		getCity(location) + getCurrentWeather(user.Lang, user.Units, f)
}

func WeatherOfDay(user *model.DB) string {
	f := forecast.GetForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeatherByDay(user, f.Daily.Data[0], f.Timezone)
}

func TomorrowWeather(user *model.DB) string {
	f := forecast.GetForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeatherByDay(user, f.Daily.Data[1], f.Timezone)
}

func WeekWeather(user *model.DB) string {
	f := forecast.GetForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeekWeather(user, f)
}
