package openweathermap

import (
	owm "github.com/briandowns/openweathermap"

	"twb/pkg/forecast"
)

type OpenWeatherMap struct {
	token string
}

// func getWeatherByDay(user *model.DB, f forecast.DataPoint, timezone string) string {
// 	return getDate(f.Time, timezone, user.Lang) + "," + getCity(user.Location) +
// 		"\n`" + f.Summary + "`\n\n" + model.Icons[f.Icon] + " *" +
// 		format.FTS0(f.TemperatureMin) + ".." + format.FTS0(f.TemperatureMax) + getTempUnit(user.Units) + "*" +
// 		"  *" + getWind(f.WindSpeed, f.WindBearing, user.Lang, user.Units) +
// 		"* \n" + model.Sunrise + " " + getTime(f.SunriseTime, timezone) +
// 		"  " + model.Sunset + " " + getTime(f.SunsetTime, timezone) +
// 		"  " + model.Moons[getMoonPhase(f.MoonPhase)] + "\n" +
// 		"`" + language.Language[user.Lang]["IFL"] + "`  *" +
// 		format.FTS0(f.ApparentTemperatureMin) + ".." + format.FTS0(f.ApparentTemperatureMax) + getTempUnit(user.Units) + "*"
// }

func (f *OpenWeatherMap) GetNow() (*forecast.Data, error) {
	data, err := owm.NewCurrent("C", "EN", f.token)
	if err != nil {
		return nil, err
	}
	if err := data.CurrentByCoordinates(&owm.Coordinates{Latitude: 59.9375, Longitude: 30.308611}); err != nil {
		return nil, err
	}
	return nil, nil
}

func New(token string) *OpenWeatherMap {
	return &OpenWeatherMap{
		token: token,
	}
}
