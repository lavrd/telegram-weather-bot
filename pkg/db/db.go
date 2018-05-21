package db

import (
	"github.com/mlbright/forecast/v2"
	"github.com/spacelavr/telegram-weather-bot/pkg/config"
	"github.com/spacelavr/telegram-weather-bot/pkg/model"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/errors"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/geocoding"
	"googlemaps.github.io/maps"
	r "gopkg.in/gorethink/gorethink.v4"
)

const (
	db    = "telegram"
	table = "users"
)

var (
	session *r.Session
)

func Init() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  config.Viper.Database.Endpoint,
		Database: db,
	})
	errors.Check(err)

	isTableAndDB()
}

func UpdateUserLang(user *model.DB, lang string, telegramID int64) string {
	ID := getUserID(telegramID)

	if user.Lang == lang {
		return lang
	} else if user.Location != "" {
		g, err := geocoding.Geocode(user.Location, lang)
		errors.Check(err)

		var data = map[string]interface{}{
			"lang":     lang,
			"location": g[0].FormattedAddress,
		}

		_, err = r.Table(table).Get(ID).Update(data).RunWrite(session)
		errors.Check(err)

		return lang
	} else {
		var data = map[string]interface{}{
			"lang": lang,
		}

		_, err := r.Table(table).Get(ID).Update(data).RunWrite(session)
		errors.Check(err)

		return lang
	}
}

func UpdateUserUnits(telegramID int64, units string) {
	if units == "°c, mps" || units == "°c, м/c" {
		units = string(forecast.SI)
	} else {
		units = string(forecast.US)
	}

	var data = map[string]interface{}{
		"units": units,
	}

	_, err := r.Table(table).Get(getUserID(telegramID)).Update(data).RunWrite(session)
	errors.Check(err)
}

func updateUserLocation(ID string, g []maps.GeocodingResult) {
	var data = map[string]interface{}{
		"location": g[0].FormattedAddress,
		"lat":      g[0].Geometry.Location.Lat,
		"lng":      g[0].Geometry.Location.Lng,
	}

	_, err := r.Table(table).Get(ID).Update(data).RunWrite(session)
	errors.Check(err)
}

func SetUser(telegramID int64, g []maps.GeocodingResult, lang string) {
	userID := getUserID(telegramID)
	if userID != nil {
		updateUserLocation(*userID, g)
		return
	}

	var data = map[string]interface{}{}
	data = map[string]interface{}{
		"telegramID": telegramID,
		"lang":       lang,
		"units":      forecast.SI,
	}

	_, err := r.Table(table).Insert(data).RunWrite(session)
	errors.Check(err)
}

func IsAuth(telegramID int64) (bool, *model.DB) {
	user := getUser(telegramID)
	if user == nil {
		return false, nil
	}
	return true, user
}
