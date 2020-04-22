package update

import (
	"fmt"
	"strings"

	"twb/pkg/emoji"
	"twb/pkg/language"
	"twb/pkg/message"
	"twb/pkg/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

const markdown = "markdown"

type MsgType int

const (
	AnyMsg MsgType = iota
	StartMsg
	BackMsg
	langKeyboardMsg
	HelpMsg
	WeatherFromCmd
	UpdateLangMsg
)

// TODO: can we throw error in this functions?

// func (m *Message) Settings(telegramID int64) {
// 	logger := prepareLogger(telegramID, "settings")

// 	_, err := m.storage.GetUser(telegramID)
// 	if err == storage.ErrUserNotFound {
// 		// LangKeyboardMsg(bot, telegramID)
// 		return
// 	}
// 	if err != nil {
// 		logger.Err(err).Msg("failed to get user")
// 		return
// 	}

// 	msg := tgbotapi.NewMessage(telegramID, model.Gear)
// 	msg.ReplyMarkup = settingsKeyboard()
// 	if _, err := m.tgBotClient.Send(msg); err != nil {
// 		logger.Err(err).Msg("failed to send message")
// 	}
// }

// func (m *Message) Units(telegramID int64) {
// 	logger := prepareLogger(telegramID, "units")

// 	user, err := m.storage.GetUser(telegramID)
// 	if err == storage.ErrUserNotFound {
// 		// LangKeyboardMsg(bot, telegramID)
// 		return
// 	}
// 	if err != nil {
// 		logger.Err(err).Msg("failed to get user")
// 		return
// 	}

// 	msg := tgbotapi.NewMessage(telegramID, model.TriangularRuler)
// 	msg.ReplyMarkup = unitsKeyboard(user.Lang)
// 	if _, err := m.tgBotClient.Send(msg); err != nil {
// 		logger.Err(err).Msg("failed to send message")
// 	}
// }

// func (m *Message) UpdateUnits(telegramID int64, units string) {
// 	logger := prepareLogger(telegramID, "update units")

// 	user, err := m.storage.GetUser(telegramID)
// 	if err == storage.ErrUserNotFound {
// 		// LangKeyboardMsg(bot, telegramID)
// 		return
// 	}
// 	if err != nil {
// 		logger.Err(err).Msg("failed to get user")
// 		return
// 	}

// 	if err := m.storage.UpdateUserUnits(telegramID, units); err != nil {
// 		logger.Err(err).Msg("failed to update user units")
// 		return
// 	}

// 	units = strings.Replace(units, units[2:3], strings.ToUpper(units[2:3]), 1)

// 	msg := tgbotapi.NewMessage(
// 		// TODO: changeUnits -> const
// 		telegramID, language.Languages[user.Lang]["changeUnits"]+" *"+units+"*",
// 	)
// 	msg.ReplyMarkup = mainKeyboard(user.Lang)
// 	msg.ParseMode = "markdown" // TODO: to const
// 	if _, err := m.tgBotClient.Send(msg); err != nil {
// 		logger.Err(err).Msg("failed to send message")
// 	}
// }

func (u *Update) MainMenuMsg(telegramID int64) {
	logger := prepareLogger(telegramID, "main menu")

	user, err := u.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		u.langKeyboardMsg(telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	msg := tgbotapi.NewMessage(telegramID, language.Dictionary[user.Lang][message.MainMenu])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	if _, err := u.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

// func (m *Message) Info(telegramID int64) {
// 	logger := prepareLogger(telegramID, "info")

// 	user, err := m.storage.GetUser(telegramID)
// 	if err == storage.ErrUserNotFound {
// 		// LangKeyboardMsg(bot, telegramID)
// 		return
// 	}
// 	if err != nil {
// 		logger.Err(err).Msg("failed to get user")
// 		return
// 	}

// 	var msg tgbotapi.MessageConfig
// 	if user.Location == "" {
// 		msg = tgbotapi.NewMessage(telegramID,
// 			"*"+language.Languages[user.Lang]["YourLLU"]+"*\n"+"`"+
// 				language.Languages[user.Lang]["empty_location"]+"`   "+
// 				model.CountriesFATE[user.Lang]+"   *"+user.Units+"*")
// 	} else {
// 		msg = tgbotapi.NewMessage(telegramID,
// 			"*"+language.Languages[user.Lang]["YourLLU"]+"*\n"+"`"+
// 				user.Location+"`   "+
// 				model.CountriesFATE[user.Lang]+"   *"+
// 				user.Units+"*")
// 	}

// 	msg.ReplyMarkup = mainKeyboard(user.Lang)
// 	msg.ParseMode = "markdown" // TODO: to const
// 	if _, err := m.tgBotClient.Send(msg); err != nil {
// 		logger.Err(err).Msg("failed to send message")
// 	}
// }

func (u *Update) UpdateLangMsg(telegramID int64, lang string) {
	logger := prepareLogger(telegramID, "update lang")

	if err := u.storage.UpdateUserLang(telegramID, emoji.CountriesFETA[lang]); err != nil {
		logger.Err(err).Msg("failed to create user")
		return
	}

	msg := tgbotapi.NewMessage(
		telegramID,
		fmt.Sprintf(
			"%s %s",
			language.Dictionary[emoji.CountriesFETA[lang]][message.ChangeLanguageTo],
			emoji.CountriesFATE[emoji.CountriesFETA[lang]],
		),
	)
	msg.ReplyMarkup = mainKeyboard(emoji.CountriesFETA[lang])
	if _, err := u.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

func (u *Update) langKeyboardMsg(telegramID int64) {
	msg := tgbotapi.NewMessage(telegramID, emoji.Globe)
	msg.ReplyMarkup = langKeyboard()
	if _, err := u.tgBotClient.Send(msg); err != nil {
		log.Err(err).Msg("failed to send message")
	}
}

func (u *Update) StartMsg(telegramID int64) {
	logger := prepareLogger(telegramID, "start")

	_, err := u.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		u.langKeyboardMsg(telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	u.helpMsg(telegramID)
}

func (u *Update) helpMsg(telegramID int64) {
	logger := prepareLogger(telegramID, "help")

	user, err := u.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		u.langKeyboardMsg(telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	msg := tgbotapi.NewMessage(telegramID, language.Dictionary[user.Lang][message.Help])
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = markdown
	if _, err := u.tgBotClient.Send(msg); err != nil {
		log.Err(err).Msg("failed to send message")
	}
}

func (u *Update) WeatherMsgByCity(telegramID int64, location string) {
	logger := prepareLogger(telegramID, "weather message by city")

	user, err := u.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		u.langKeyboardMsg(telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user") // TODO: Return message about error everywhere.
		return
	}

	geoRes, err := u.geocode.Geocode(location, user.Lang)
	if err != nil {
		logger.Err(err).Msg("failed to geocode")
		return
	}

	// if g, err := geocoding.Geocode(location, user.Lang); err != nil {
	// 	msg = tgbotapi.NewMessage(telegramID, err.Error())
	// } else {
	// 	if user.Location != g[0].FormattedAddress {
	// 		db.SetUser(telegramID, g, user.Lang)
	//
	// 		msg := tgbotapi.NewMessage(telegramID,
	// 			language.Language[user.Lang]["changeCityTo"]+" "+g[0].FormattedAddress)
	// 		_, err = bot.Send(msg)
	// 		errors.Check(err)
	// 	}
	//
	// 	wthr := weather.CurrentWeather(
	// 		g[0].Geometry.Location.Lat, g[0].Geometry.Location.Lng,
	// 		g[0].FormattedAddress, user)
	//
	// 	msg = tgbotapi.NewMessage(telegramID, wthr)
	// }
	//

	msg := tgbotapi.NewMessage(telegramID, fmt.Sprintf("%v %v %v", geoRes.Location, geoRes.Lat, geoRes.Lon))
	msg.ReplyMarkup = mainKeyboard(user.Lang)
	msg.ParseMode = "markdown"
	if _, err := u.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}

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

func (u *Update) weatherMsgByCmd(telegramID int64, weatherType string) {
	logger := prepareLogger(telegramID, "weather message by command")

	user, err := u.storage.GetUser(telegramID)
	if err == storage.ErrUserNotFound {
		u.langKeyboardMsg(telegramID)
		return
	}
	if err != nil {
		logger.Err(err).Msg("failed to get user")
		return
	}

	var (
		msg     tgbotapi.MessageConfig
		weather string
	)

	if user.Location == "" {
		msg = tgbotapi.NewMessage(telegramID, language.Dictionary[user.Lang][message.EmptyCity])
		msg.ReplyMarkup = helpKeyboard()
	} else {
		switch weatherType {
		case strings.ToLower(language.Dictionary[language.EN][message.Now]), strings.ToLower(language.Dictionary[language.RU][message.Now]):
			// 	wthr = weather.CurrentWeather(user.Lat, user.Lng, user.Location, user)

		case strings.ToLower(language.Dictionary[language.EN][message.ForToday]), strings.ToLower(language.Dictionary[language.RU][message.ForToday]):
			// 	wthr = weather.WeatherOfDay(user)

		case strings.ToLower(language.Dictionary[language.EN][message.ForTomorrow]), strings.ToLower(language.Dictionary[language.RU][message.ForTomorrow]):
			// 	wthr = weather.TomorrowWeather(user)

		case strings.ToLower(language.Dictionary[language.EN][message.ForWeek]), strings.ToLower(language.Dictionary[language.RU][message.ForWeek]):
			// 	wthr = weather.WeekWeather(user)
		}

		msg = tgbotapi.NewMessage(telegramID, weather)
		msg.ReplyMarkup = mainKeyboard(user.Lang)
		msg.ParseMode = markdown
	}

	if _, err := u.tgBotClient.Send(msg); err != nil {
		logger.Err(err).Msg("failed to send message")
	}
}
