package db

import (
	m "github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"github.com/lavrs/telegram-weather-bot/utils/geocoding"
	"github.com/mlbright/forecast/v2"
	"googlemaps.github.io/maps"
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
)

const (
	db    = "telegram"
	table = "users"
)

var (
	session *r.Session
	err     error
)

func init() {
	// open db session (connect to db)
	if session, err = r.Connect(r.ConnectOpts{
		Address:  "172.17.0.2:28015",
		Database: db,
	}); err != nil {
		log.Panic(err)
	}

	// check if db and table exists
	isTableAndDB()
}

// UpdateUserLang update user lang
func UpdateUserLang(user *m.DB, lang string, telegramID int64) string {
	ID := getUserID(telegramID)

	// if user lang == lang, no need to change lang
	if user.Lang == lang {
		return lang
	} else if user.Location != "" {
		// if user location is exists, translate location

		// geocoding in another language
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
		// just change language if user location is empty

		var data = map[string]interface{}{
			"lang": lang,
		}

		_, err := r.Table(table).Get(ID).Update(data).RunWrite(session)
		errors.Check(err)
		return lang
	}
}

// UpdateUserUnits update user units
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

// update user location
func updateUserLocation(ID string, g []maps.GeocodingResult) {
	var data = map[string]interface{}{
		"location": g[0].FormattedAddress,
		"lat":      g[0].Geometry.Location.Lat,
		"lng":      g[0].Geometry.Location.Lng,
	}

	_, err := r.Table(table).Get(ID).Update(data).RunWrite(session)
	errors.Check(err)
}

// SetUser set new user
func SetUser(telegramID int64, g []maps.GeocodingResult, lang string) {
	userID := getUserID(telegramID)

	// if user already exists, update user
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

// IsAuth check auth
// if user auth, returns user info
func IsAuth(telegramID int64) (bool, *m.DB) {
	user := getUser(telegramID)

	if user == nil {
		return false, nil
	}
	return true, user
}
