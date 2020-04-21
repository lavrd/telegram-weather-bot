package openweathermap

import (
	owm "github.com/briandowns/openweathermap"

	"twb/pkg/forecast"
)

type OpenWeatherMap struct {
	token string
}

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
