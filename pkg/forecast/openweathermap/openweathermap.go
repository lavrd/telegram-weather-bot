package openweathermap

import (
	"net/http"

	"twb/pkg/forecast"
	"twb/pkg/utils/converter"
)

type OpenWeatherMap struct {
	token      string
	httpClient *http.Client
}

func (owm *OpenWeatherMap) GetNow() (*forecast.Current, error) {
	data, err := owm.req("55.751244", "37.618423", "imperial", "RU")
	if err != nil {
		return nil, err
	}

	forecastData := &forecast.Current{
		Temp:         converter.FTS0(data.Current.Temp),
		ApparentTemp: converter.FTS0(data.Current.FeelsLike),
		Wind: &forecast.Wind{
			Speed: 0,
			Deg:   0,
		},
	}

	for _, weather := range data.Current.Weather {
		forecastData.Conditions = append(forecastData.Conditions, &forecast.Condition{
			Summary: weather.Description,
			Type:    parseCondition(weather.ID),
		})
	}

	return forecastData, nil
}

func New(token string) *OpenWeatherMap {
	return &OpenWeatherMap{
		token:      token,
		httpClient: &http.Client{},
	}
}
