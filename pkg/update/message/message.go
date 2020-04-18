package message

import (
	"strings"
	"telegram-weather-bot/pkg/language"
	"telegram-weather-bot/pkg/model"
	"telegram-weather-bot/pkg/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// type Type int

// const (
// 	WeatherCmd Type = iota + 1
// )

type Message struct {
	storage storage.Storage

	tgBotClient *tgbotapi.BotAPI
}

// TODO: can we throw error in this functions?

func (m *Message) Settings(telegramID int64) {
	logger := prepareLogger(telegramID, "settings")

	_, err := m.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		// LangKeyboardMsg(bot, telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	msg := tgbotapi.NewMessage(telegramID, model.Gear)
	msg.ReplyMarkup = settingsKeyboard()
	if _, err := m.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

func (m *Message) Units(telegramID int64) {
	logger := prepareLogger(telegramID, "units")

	user, err := m.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		// LangKeyboardMsg(bot, telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	msg := tgbotapi.NewMessage(telegramID, model.TriangularRuler)
	msg.ReplyMarkup = unitsKeyboard(user.Lang)
	if _, err := m.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

func (m *Message) UpdateUnits(telegramID int64, units string) {
	logger := prepareLogger(telegramID, "update units")

	user, err := m.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		// LangKeyboardMsg(bot, telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	if err := m.storage.UpdateUserUnits(telegramID, units); err != nil {
		logger.Err(err).Msg("failed to update user units")
		return
	}

	units = strings.Replace(units, units[2:3], strings.ToUpper(units[2:3]), 1)

	msg := tgbotapi.NewMessage(
		// TODO: changeUnits -> const
		telegramID, language.Languages[user.Lang]["changeUnits"]+" *"+units+"*",
	)
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown" // TODO: to const
	if _, err := m.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

func (m *Message) MainMenu(telegramID int64) {
	logger := prepareLogger(telegramID, "main menu")

	user, err := m.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		// LangKeyboardMsg(bot, telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	msg := tgbotapi.NewMessage(telegramID, language.Languages[user.Lang]["mainMenu"]) // TODO: const
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	if _, err := m.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

func (m *Message) Info(telegramID int64) {
	logger := prepareLogger(telegramID, "info")

	user, err := m.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		// LangKeyboardMsg(bot, telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	var msg tgbotapi.MessageConfig
	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID,
			"*"+language.Languages[user.Lang]["YourLLU"]+"*\n"+"`"+
				language.Languages[user.Lang]["empty_location"]+"`   "+
				model.CountriesFATE[user.Lang]+"   *"+user.Units+"*")
	} else {
		msg = tgbotapi.NewMessage(telegramID,
			"*"+language.Languages[user.Lang]["YourLLU"]+"*\n"+"`"+
				user.Location+"`   "+
				model.CountriesFATE[user.Lang]+"   *"+
				user.Units+"*")
	}

	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown" // TODO: to const
	if _, err := m.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

// func UpdateLangMsg(bot *tgbotapi.BotAPI, telegramID int64, message string) {
// 	isAuth, user := db.IsAuth(telegramID)
// 	var msg tgbotapi.MessageConfig

// 	if isAuth {
// 		lang := db.UpdateUserLang(user, model.CountriesFETA[message], telegramID)

// 		msg = tgbotapi.NewMessage(telegramID,
// 			language.Language[lang]["changeLanguageTo"]+" "+model.CountriesFATE[model.CountriesFETA[message]])
// 		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
// 	} else {
// 		db.SetUser(telegramID, nil, model.CountriesFETA[message])

// 		msg = tgbotapi.NewMessage(telegramID,
// 			language.Language[model.CountriesFETA[message]]["changeLanguageTo"]+" "+
// 				model.CountriesFATE[model.CountriesFETA[message]])
// 		msg.ReplyMarkup = mainKeyboard(model.CountriesFETA[message])
// 	}

// 	_, err := bot.Send(msg)
// 	errors.Check(err)
// }

// func LangKeyboardMsg(bot *tgbotapi.BotAPI, telegramID int64) {
// 	msg := tgbotapi.NewMessage(telegramID, model.GlobeWithMeridian)
// 	msg.ReplyMarkup = langKeyboard()
// 	_, err := bot.Send(msg)
// 	errors.Check(err)
// }

// func StartMsg(bot *tgbotapi.BotAPI, telegramID int64) {
// 	isAuth, _ := db.IsAuth(telegramID)

// 	if !isAuth {
// 		LangKeyboardMsg(bot, telegramID)
// 		return
// 	}
// 	Help(bot, telegramID)
// }

// func Help(bot *tgbotapi.BotAPI, telegramID int64) {
// 	isAuth, user := db.IsAuth(telegramID)

// 	if !isAuth {
// 		LangKeyboardMsg(bot, telegramID)
// 		return
// 	}

// 	msg := tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["help"])
// 	msg.ReplyMarkup = mainKeyboard(user.Lang)
// 	msg.ParseMode = "markdown"
// 	_, err := bot.Send(msg)
// 	errors.Check(err)
// }

// func WeatherMsgFromCity(bot *tgbotapi.BotAPI, telegramID int64, location string) {
// 	isAuth, user := db.IsAuth(telegramID)
// 	var msg tgbotapi.MessageConfig

// 	if !isAuth {
// 		LangKeyboardMsg(bot, telegramID)
// 		return
// 	}

// 	if g, err := geocoding.Geocode(location, user.Lang); err != nil {
// 		msg = tgbotapi.NewMessage(telegramID, err.Error())
// 	} else {
// 		if user.Location != g[0].FormattedAddress {
// 			db.SetUser(telegramID, g, user.Lang)

// 			msg := tgbotapi.NewMessage(telegramID,
// 				language.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
// 			_, err = bot.Send(msg)
// 			errors.Check(err)
// 		}

// 		wthr := weather.CurrentWeather(
// 			g[0].Geometry.Location.Lat, g[0].Geometry.Location.Lng,
// 			g[0].FormattedAddress, user)

// 		msg = tgbotapi.NewMessage(telegramID, wthr)
// 	}

// 	msg.ReplyMarkup = mainKeyboard(user.Lang)
// 	msg.ParseMode = "markdown"
// 	_, err := bot.Send(msg)
// 	errors.Check(err)
// }

// func WeatherMsgFromLocation(bot *tgbotapi.BotAPI, telegramID int64, location *tgbotapi.Location) {
// 	isAuth, user := db.IsAuth(telegramID)
// 	var msg tgbotapi.MessageConfig

// 	if !isAuth {
// 		LangKeyboardMsg(bot, telegramID)
// 		return
// 	}

// 	if g, err := geocoding.ReverseGeocode(location, user.Lang); err != nil {
// 		msg = tgbotapi.NewMessage(telegramID, err.Error())
// 	} else {
// 		if user.Lat != g[0].Geometry.Location.Lat ||
// 			user.Lng != g[0].Geometry.Location.Lng {

// 			db.SetUser(telegramID, g, user.Lang)

// 			msg = tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
// 			msg.ReplyMarkup = mainKeyboard(user.Lang)
// 			_, err = bot.Send(msg)
// 			errors.Check(err)
// 		}

// 		wthr := weather.CurrentWeatherFromLocation(g[0].Geometry.Location.Lat,
// 			g[0].Geometry.Location.Lng, g[0].FormattedAddress, user)

// 		msg = tgbotapi.NewMessage(telegramID, wthr)
// 	}

// 	msg.ParseMode = "markdown"
// 	_, err := bot.Send(msg)
// 	errors.Check(err)
// }

// func WeatherMsgFromCmd(bot *tgbotapi.BotAPI, telegramID int64, message string) {
// 	isAuth, user := db.IsAuth(telegramID)

// 	var (
// 		msg  tgbotapi.MessageConfig
// 		wthr string
// 		err  error
// 	)

// 	if !isAuth {
// 		LangKeyboardMsg(bot, telegramID)
// 		return
// 	}

// 	if user.Location == "" {
// 		msg = tgbotapi.NewMessage(telegramID, language.Language[user.Lang]["emptycity"])
// 		msg.ReplyMarkup = helpKeyboard()
// 	} else {
// 		switch {
// 		case message == "now" || message == "/now" || message == "сейчас":
// 			wthr = weather.CurrentWeather(user.Lat, user.Lng, user.Location, user)

// 		case message == "for today" || message == "/today" || message == "на сегодня":
// 			wthr = weather.WeatherOfDay(user)

// 		case message == "for tomorrow" || message == "/tomorrow" || message == "на завтра":
// 			wthr = weather.TomorrowWeather(user)

// 		case message == "for week" || message == "/week" || message == "на неделю":
// 			wthr = weather.WeekWeather(user)
// 		}

// 		msg = tgbotapi.NewMessage(telegramID, wthr)
// 		msg.ReplyMarkup = mainKeyboard(user.Lang)
// 		msg.ParseMode = "markdown"
// 	}

// 	_, err = bot.Send(msg)
// 	errors.Check(err)
// }
