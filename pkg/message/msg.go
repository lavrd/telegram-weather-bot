package msg

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spacelavr/telegram-weather-bot/pkg/db"
	"github.com/spacelavr/telegram-weather-bot/pkg/language"
	"github.com/spacelavr/telegram-weather-bot/pkg/model"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/errors"
	"github.com/spacelavr/telegram-weather-bot/pkg/utils/geocoding"
	"github.com/spacelavr/telegram-weather-bot/pkg/weather"
)

func SettingsMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, _ := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, model.Gear)
	msg.ReplyMarkup = settingsKeyboard()
	_, err := bot.Send(msg)
	errors.Check(err)
}

func UnitsMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, model.TriangularRuler)
	msg.ReplyMarkup = unitsKeyboard(user.Lang)
	_, err := bot.Send(msg)
	errors.Check(err)
}

func UpdateUnitsMsg(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	db.UpdateUserUnits(telegramID, message)

	message = strings.Replace(message, message[2:3],
		strings.ToUpper(message[2:3]), 1)

	msg := tgbotapi.NewMessage(telegramID,
		language.Language[user.Lang]["changeUnits"]+" *"+message+"*")
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

func MainMenuMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["mainMenu"])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	_, err := bot.Send(msg)
	errors.Check(err)
}

func InfoMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID,
			"*" + language.Language[user.Lang]["YourLLU"] + "*\n" + "`"+
				language.Language[user.Lang]["empty_location"]+ "`   "+
				model.CountriesFATE[user.Lang]+ "   *"+ user.Units+ "*")
	} else {
		msg = tgbotapi.NewMessage(telegramID,
			"*" + language.Language[user.Lang]["YourLLU"] + "*\n" + "`"+
				user.Location+ "`   "+
				model.CountriesFATE[user.Lang]+ "   *"+
				user.Units+ "*")
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

func UpdateLangMsg(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	if isAuth {
		lang := db.UpdateUserLang(user, model.CountriesFETA[message], telegramID)

		msg = tgbotapi.NewMessage(telegramID,
			language.Language[lang]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	} else {
		db.SetUser(telegramID, nil, model.CountriesFETA[message])

		msg = tgbotapi.NewMessage(telegramID,
			language.Language[model.CountriesFETA[message]]["changeLanguageTo"] + " "+
				model.CountriesFATE[model.CountriesFETA[message]])
		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
	}

	_, err := bot.Send(msg)
	errors.Check(err)
}

func LangKeyboardMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	msg := tgbotapi.NewMessage(telegramID, model.GlobeWithMeridian)
	msg.ReplyMarkup = langKeyboard()
	_, err := bot.Send(msg)
	errors.Check(err)
}

func StartMsg(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, _ := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}
	Help(bot, telegramID)
}

func Help(bot *tgbotapi.BotAPI, telegramID int64) {
	isAuth, user := db.IsAuth(telegramID)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	msg := tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["help"])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

func WeatherMsgFromCity(bot *tgbotapi.BotAPI, telegramID int64, location string) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	if g, err := geocoding.Geocode(location, user.Lang); err != nil {
		msg = tgbotapi.NewMessage(telegramID, err.Error())
	} else {
		if user.Location != g[0].FormattedAddress {
			db.SetUser(telegramID, g, user.Lang)

			msg := tgbotapi.NewMessage(telegramID,
				language.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
			_, err = bot.Send(msg)
			errors.Check(err)
		}

		wthr := weather.CurrentWeather(
			g[0].Geometry.Location.Lat, g[0].Geometry.Location.Lng,
			g[0].FormattedAddress, user)

		msg = tgbotapi.NewMessage(telegramID, wthr)
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

func WeatherMsgFromLocation(bot *tgbotapi.BotAPI, telegramID int64, location *tgbotapi.Location) {
	isAuth, user := db.IsAuth(telegramID)
	var msg tgbotapi.MessageConfig

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	if g, err := geocoding.ReverseGeocode(location, user.Lang); err != nil {
		msg = tgbotapi.NewMessage(telegramID, err.Error())
	} else {
		if user.Lat != g[0].Geometry.Location.Lat ||
			user.Lng != g[0].Geometry.Location.Lng {

			db.SetUser(telegramID, g, user.Lang)

			msg = tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
			msg.ReplyMarkup = mainKeyboard(user.Lang)
			_, err = bot.Send(msg)
			errors.Check(err)
		}

		wthr := weather.CurrentWeatherFromLocation(g[0].Geometry.Location.Lat,
			g[0].Geometry.Location.Lng, g[0].FormattedAddress, user)

		msg = tgbotapi.NewMessage(telegramID, wthr)
	}

	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	errors.Check(err)
}

func WeatherMsgFromCmd(bot *tgbotapi.BotAPI, telegramID int64, message string) {
	isAuth, user := db.IsAuth(telegramID)

	var (
		msg  tgbotapi.MessageConfig
		wthr string
		err  error
	)

	if !isAuth {
		LangKeyboardMsg(bot, telegramID)
		return
	}

	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["emptycity"])
		msg.ReplyMarkup = helpKeyboard()
	} else {
		switch {
		case message == "now" || message == "/now" || message == "сейчас":
			wthr = weather.CurrentWeather(user.Lat, user.Lng, user.Location, user)

		case message == "for today" || message == "/today" || message == "на сегодня":
			wthr = weather.WeatherOfDay(user)

		case message == "for tomorrow" || message == "/tomorrow" || message == "на завтра":
			wthr = weather.TomorrowWeather(user)

		case message == "for week" || message == "/week" || message == "на неделю":
			wthr = weather.WeekWeather(user)
		}

		msg = tgbotapi.NewMessage(telegramID, wthr)
		msg.ReplyMarkup = mainKeyboard(user.Lang)
		msg.ParseMode = "markdown"
	}

	_, err = bot.Send(msg)
	errors.Check(err)
}
