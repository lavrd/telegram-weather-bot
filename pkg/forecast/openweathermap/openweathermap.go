package openweathermap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"

	"twb/pkg/forecast"
	"twb/pkg/utils/converter"
)

const baseURLFormat = "https://api.openweathermap.org/data/2.5/onecall?lat=%f&lon=%f&units=%s&lang=%s&appid=%s"

var ErrUnknown = errors.New("unknown error")

type Data struct {
	Current *Current `json:"current"`
}

type Current struct {
	Temp      float64    `json:"temp"`
	FeelsLike float64    `json:"feels_like"`
	WindSpeed float64    `json:"wind_speed"`
	WindDeg   int        `json:"wind_deg"`
	Weather   []*Weather `json:"weather"`
}

type Weather struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type OpenWeatherMap struct {
	token      string
	httpClient *http.Client
}

func (owm *OpenWeatherMap) GetCurrent(lat, lon float64, units, lang string) (*forecast.Current, error) {
	data, err := owm.req(lat, lon, units, lang)
	if err != nil {
		return nil, err
	}

	forecastData := &forecast.Current{
		Temp:         converter.FTS0(data.Current.Temp),
		ApparentTemp: converter.FTS0(data.Current.FeelsLike),
		Wind: &forecast.Wind{
			Speed: converter.FTS0(data.Current.WindSpeed),
			Deg:   data.Current.WindDeg,
		},
	}

	for _, weather := range data.Current.Weather {
		forecastData.Conditions = append(forecastData.Conditions, &forecast.Condition{
			Summary: weather.Description,
			Type:    parseConditionByID(weather.ID),
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

func (owm *OpenWeatherMap) req(lat, lon float64, units, lang string) (*Data, error) {
	url := fmt.Sprintf(baseURLFormat, lat, lon, units, lang, owm.token)
	res, err := owm.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Error().Int("status code", res.StatusCode).Msg("failed to get weather: invalid response status code")
		return nil, ErrUnknown
	}

	data := &Data{}
	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

func parseConditionByID(id int) forecast.ConditionType {
	switch {
	case id >= 200 && id <= 232:
		return forecast.Thunderstorm
	case id >= 300 && id <= 321:
		return forecast.Drizzle
	case id >= 500 && id <= 531:
		return forecast.Rain
	case id >= 600 && id <= 622:
		return forecast.Snow
	case id == 701:
		return forecast.Mist
	case id == 711:
		return forecast.Smoke
	case id == 721:
		return forecast.Haze
	case id == 731:
		return forecast.Dust
	case id == 741:
		return forecast.Fog
	case id == 751:
		return forecast.Sand
	case id == 761:
		return forecast.Dust
	case id == 762:
		return forecast.Ash
	case id == 771:
		return forecast.Squall
	case id == 781:
		return forecast.Tornado
	case id == 800:
		return forecast.Clear
	case id >= 801 && id <= 804:
		return forecast.Clouds
	default:
		return forecast.Unknown
	}
}
