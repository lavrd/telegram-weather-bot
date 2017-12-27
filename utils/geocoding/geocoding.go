package geocoding

import (
	"context"
	e "errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	c "github.com/lavrs/telegram-weather-bot/config"
	l "github.com/lavrs/telegram-weather-bot/language"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"googlemaps.github.io/maps"
)

// ReverseGeocode geocoding location to location name
func ReverseGeocode(location *tgbotapi.Location, lang string) ([]maps.GeocodingResult, error) {
	var g []maps.GeocodingResult

	client, err := maps.NewClient(maps.WithAPIKey(c.Cfg.GoogleGeocodingToken))
	errors.Check(err)

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
			return nil, e.New("_" + l.Language[lang]["ZERO_RESULTS_LOCATION"] + "_")
		} else {
			return nil, e.New("_" + l.Language[lang]["unknownError"] + "_")
		}
	}

	return g, nil
}

// Geocode geocoding location name to location
func Geocode(location, lang string) ([]maps.GeocodingResult, error) {
	var g []maps.GeocodingResult

	client, err := maps.NewClient(maps.WithAPIKey(c.Cfg.GoogleGeocodingToken))
	errors.Check(err)

	r := &maps.GeocodingRequest{
		Address:  location,
		Language: lang,
	}

	if g, err = client.Geocode(context.Background(), r); err != nil {
		if err.Error() == "maps: ZERO_RESULTS - " {
			return nil, e.New("_" + l.Language[lang]["ZERO_RESULTS_CITY"] + "_")
		} else {
			return nil, e.New("_" + l.Language[lang]["unknownError"] + "_")
		}
	}

	return g, nil
}
