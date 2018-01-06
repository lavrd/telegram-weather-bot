package msg

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spacelavr/telegram-weather-bot/db"
	l "github.com/spacelavr/telegram-weather-bot/language"
	"github.com/spacelavr/telegram-weather-bot/model"
	"github.com/spacelavr/telegram-weather-bot/utils/errors"
	"github.com/spacelavr/telegram-weather-bot/utils/geocoding"
	w "github.com/spacelavr/telegram-weather-bot/weather"
)

// SettingsMsg send settings message
func SettingsMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, _ := db.IsAuth(telegramID)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, model.Gear)
	msg.ReplyMarkup = settingsKeyboard()
	_, err := bot.Send(msg)
	errors.Check(err)
}

// UnitsMsg send units message
func UnitsMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, model.TriangularRuler)
	msg.ReplyMarkup = unitsKeyboard(user.Lang)
	_, err := bot.Send(msg)
	errors.Check(err)
}

// UpdateUnitsMsg send complete update units message
func UpdateUnitsMsg(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	db.UpdateUserUnits(telegramID, message)

	message = strings.Replace(message, message[2:3],
		strings.ToUpper(message[2:3]), 1)

	msg := tgbotapi.NewMessage(telegramID,
		l.Language[user.Lang]["changeUnits"]+" *"+message+"*")
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

// MainMenuMsg return to main menu
func MainMenuMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["mainMenu"])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	_, err := bot.Send(msg)
	errors.Check(err)
}

func InfoMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	// if user location empty
	// send message with empty location
	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID,
			"*" + l.Language[user.Lang]["YourLLU"] + "*\n" + "`"+
				l.Language[user.Lang]["empty_location"]+ "`   "+
				model.CountriesFATE[user.Lang]+ "   *"+ user.Units+ "*")
	} else {
		// send with user location
		msg = tgbotapi.NewMessage(telegramID,
			"*" + l.Language[user.Lang]["YourLLU"] + "*\n" + "`"+
				user.Location+ "`   "+
				model.CountriesFATE[user.Lang]+ "   *"+
				user.Units+ "*")
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

// UpdateLangMsg send complete update user language
func UpdateLangMsg(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	// if user existts
	// update user language
	if isAuth {
		lang := db.UpdateUserLang(user, model.CountriesFETA[message], telegramID)

		msg = tgbotapi.NewMessage(telegramID,
			l.Language[lang]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	} else {
		// if user not exists
		// added new user
		db.SetUser(telegramID, nil, model.CountriesFETA[message])

		msg = tgbotapi.NewMessage(telegramID,
			l.Language[model.CountriesFETA[message]]["changeLanguageTo"] + " "+
				model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	}

	_, err := bot.Send(msg)
	errors.Check(err)
}

// LangKeyboardMsg show lang keyboard
func LangKeyboardMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	msg := tgbotapi.NewMessage(telegramID, model.GlobeWithMeridian)
	msg.ReplyMarkup = langKeyboard()
	_, err := bot.Send(msg)
	errors.Check(err)
}

// StartMsg show lang keyboard at start
// if user exists show help message
func StartMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, _ := db.IsAuth(telegramID)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}
	Help(bot, telegramID)
}

// Help show help message
func Help(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["help"])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

// WeatherMsgFromCity show weather by city name
func WeatherMsgFromCity(bot *tgbotapi.BotAPI, telegramID int64, location string) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	// geocoding city name
	if g, err := geocoding.Geocode(location, user.Lang); err != nil {
		// if geocoding error, send to user
		msg = tgbotapi.NewMessage(telegramID, err.Error())
	} else {
		// send weather
		// if user location != city name, set new user location
		if user.Location != g[0].FormattedAddress {
			db.SetUser(telegramID, g, user.Lang)

			// send complete change city message
			msg := tgbotapi.NewMessage(telegramID,
				l.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
			_, err = bot.Send(msg)
			errors.Check(err)
		}

		// get weather
		wthr := w.CurrentWeather(
			g[0].Geometry.Location.Lat, g[0].Geometry.Location.Lng,
			g[0].FormattedAddress, user)

		msg = tgbotapi.NewMessage(telegramID, wthr)
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

// WeatherMsgFromLocation show weather by location
func WeatherMsgFromLocation(bot *tgbotapi.BotAPI, telegramID int64, location *tgbotapi.Location) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	// reverse geocoding location (to city name)
	if g, err := geocoding.ReverseGeocode(location, user.Lang); err != nil {
		// if geocoding error, send to user
		msg = tgbotapi.NewMessage(telegramID, err.Error())
	} else {
		// if user location != current location, set new user location
		if user.Lat != g[0].Geometry.Location.Lat ||
			user.Lng != g[0].Geometry.Location.Lng {

			db.SetUser(telegramID, g, user.Lang)

			msg = tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
			msg.ReplyMarkup = mainKeyboard(user.Lang)
			_, err = bot.Send(msg)
			errors.Check(err)
		}

		// get weather
		wthr := w.CurrentWeatherFromLocation(g[0].Geometry.Location.Lat,
			g[0].Geometry.Location.Lng, g[0].FormattedAddress, user)

		msg = tgbotapi.NewMessage(telegramID, wthr)
	}

	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

// WeatherMsgFromCmd show weather by user location for now, today, tomorrow, week
func WeatherMsgFromCmd(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)

	var (
		msg  tgbotapi.MessageConfig
		wthr string
		err  error
	)

	// if user not auth (exists), send lang keyboard
	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	// if user location is empty
	// send to user that user location is empty
	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID, l.Language[user.Lang]["emptycity"])
		msg.ReplyMarkup = helpKeyboard()
	} else {
		// choose weather time
		switch {
		case message == "now" || message == "/now" || message == "сейчас":
			wthr = w.CurrentWeather(user.Lat, user.Lng, user.Location, user)

		case message == "for today" || message == "/today" || message == "на сегодня":
			wthr = w.WeatherOfDay(user)

		case message == "for tomorrow" || message == "/tomorrow" || message == "на завтра":
			wthr = w.TomorrowWeather(user)

		case message == "for week" || message == "/week" || message == "на неделю":
			wthr = w.WeekWeather(user)
		}

		msg = tgbotapi.NewMessage(telegramID, wthr)
		msg.ReplyMarkup = mainKeyboard(user.Lang)
		msg.ParseMode = "markdown"
	}

	_, err = bot.Send(msg)
	errors.Check(err)
}
