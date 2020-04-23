package google

import (
	"context"

	"twb/pkg/geocode"

	"googlemaps.github.io/maps"
)

type Google struct {
	client *maps.Client
}

func (g *Google) Geocode(location, lang string) (*geocode.Result, error) {
	req := &maps.GeocodingRequest{
		Language: lang,
		Address:  location,
	}

	ress, err := g.client.Geocode(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	if len(ress) == 0 {
		return nil, geocode.ErrEmptyResult
	}
	res := ress[0]

	return &geocode.Result{
		Location: res.FormattedAddress,
		Lat:      res.Geometry.Location.Lat,
		Lon:      res.Geometry.Location.Lng,
	}, nil
}

func (g *Google) Reverse(lat, lon float64, lang string) (*geocode.Result, error) {
	req := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: lat,
			Lng: lon,
		},
		Language: lang,
	}

	ress, err := g.client.ReverseGeocode(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	res := ress[0]

	return &geocode.Result{
		Location: res.FormattedAddress,
		Lat:      res.Geometry.Location.Lat,
		Lon:      res.Geometry.Location.Lng,
	}, nil
}

func New(token string) (*Google, error) {
	client, err := maps.NewClient(maps.WithAPIKey(token))
	if err != nil {
		return nil, err
	}

	return &Google{
		client: client,
	}, nil
}
