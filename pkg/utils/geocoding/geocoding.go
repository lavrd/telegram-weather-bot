package geocoding

import (
	"context"
	"errors"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spacelavr/telegram-weather-bot/pkg/config"
	"github.com/spacelavr/telegram-weather-bot/pkg/language"
	ue "github.com/spacelavr/telegram-weather-bot/pkg/utils/errors"
	"googlemaps.github.io/maps"
)

func ReverseGeocode(location *tgbotapi.Location, lang string) ([]maps.GeocodingResult, error) {
	var g []maps.GeocodingResult

	client, err := maps.NewClient(maps.WithAPIKey(config.Viper.Google.Geocoding.Token))
	ue.Check(err)

	latLng := &maps.LatLng{
		Lat: location.Latitude,
		Lng: location.Longitude,
	}

	r := &maps.GeocodingRequest{
		LatLng:   latLng,
		Language: lang,
	}

	if g, err = client.ReverseGeocode(context.Background(), r); err != nil {
		if err.Error() == "maps: ZERO_RESULTS - " {
			return nil, errors.New("_" + language.Language[lang]["ZERO_RESULTS_LOCATION"] + "_")
		} else {
			return nil, errors.New("_" + language.Language[lang]["unknownError"] + "_")
		}
	}

	return g, nil
}

func Geocode(location, lang string) ([]maps.GeocodingResult, error) {
	var g []maps.GeocodingResult

	client, err := maps.NewClient(maps.WithAPIKey(config.Viper.Google.Geocoding.Token))
	ue.Check(err)

	r := &maps.GeocodingRequest{
		Address:  location,
		Language: lang,
	}

	if g, err = client.Geocode(context.Background(), r); err != nil {
		if err.Error() == "maps: ZERO_RESULTS - " {
			return nil, errors.New("_" + language.Language[lang]["ZERO_RESULTS_CITY"] + "_")
		} else {
			return nil, errors.New("_" + language.Language[lang]["unknownError"] + "_")
		}
	}

	return g, nil
}
