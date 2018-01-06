package weather

import (
	"github.com/spacelavr/telegram-weather-bot/model"
	"github.com/spacelavr/telegram-weather-bot/utils/forecast"
)

// CurrentWeather returns weather for now from location name
func CurrentWeather(lat, lng float64, location string, user *model.DB) string {
	f := forecast.GetForecast(lat, lng, user.Lang, user.Units)

	return getTime(f.Currently.Time, f.Timezone) +
		getCity(location) + getCurrentWeather(user.Lang, user.Units, f)
}

// CurrentWeatherFromLocation returns weather for now from location
func CurrentWeatherFromLocation(lat, lng float64, location string, user *model.DB) string {
	f := forecast.GetForecast(lat, lng, user.Lang, user.Units)

	return getTime(f.Currently.Time, f.Timezone) +
		getCity(location) + getCurrentWeather(user.Lang, user.Units, f)
}

// WeatherOfDay returns weather for today
func WeatherOfDay(user *model.DB) string {
	f := forecast.GetForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeatherByDay(user, f.Daily.Data[0], f.Timezone)
}

// TomorrowWeather returns weather for tomorrow
func TomorrowWeather(user *model.DB) string {
	f := forecast.GetForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeatherByDay(user, f.Daily.Data[1], f.Timezone)
}

// WeekWeather returns weather for week
func WeekWeather(user *model.DB) string {
	f := forecast.GetForecast(user.Lat, user.Lng, user.Lang, user.Units)

	return getWeekWeather(user, f)
}
