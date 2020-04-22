package update

import (
	"strings"

	"twb/pkg/forecast"
	"twb/pkg/geocode"
	"twb/pkg/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Update struct {
	tgBotClient *tgbotapi.BotAPI

	storage storage.Storage

	geocode  geocode.Geocode
	forecast forecast.Forecast
}

func (u *Update) Handle(upd *tgbotapi.Update) {
	if upd.Message == nil {
		return
	}

	msgText := strings.ToLower(upd.Message.Text)
	cmdText := upd.Message.Command()
	telegramID := upd.Message.Chat.ID

	msgType := parseIncomingMsg(msgText, cmdText)

	switch msgType {
	case StartMsg:
		u.StartMsg(telegramID)
	case BackMsg:
		u.MainMenuMsg(telegramID)
	case UpdateLangMsg:
		u.UpdateLangMsg(telegramID, msgText)
	case langKeyboardMsg:
		u.langKeyboardMsg(telegramID)
	case HelpMsg:
		u.helpMsg(telegramID)
	case WeatherFromCmd:
		u.weatherMsgByCmd(telegramID, msgText)
	default:
		u.WeatherMsgByCity(telegramID, msgText)
	}

	// if update.Message.Text == model.Info || update.Message.Command() == "info" {
	// 	InfoMsg(bot, update.Message.Chat.ID)

	// 	return
	// }

	// if update.Message.Text == model.Gear {
	// 	SettingsMsg(bot, update.Message.Chat.ID)

	// 	return
	// }

	// if update.Message.Text == "°c, mps" || update.Message.Text == "°c, м/c" ||
	// 	update.Message.Text == "°f, mph" || update.Message.Text == "°f, миль/ч" {

	// 	UpdateUnitsMsg(bot, update.Message.Chat.ID, update.Message.Text)

	// 	return
	// }

	// if update.Message.Text == model.TriangularRuler || update.Message.Command() == "units" {
	// 	UnitsMsg(bot, update.Message.Chat.ID)

	// 	return
	// }

	// if update.Message.Location != nil {
	// 	WeatherMsgFromLocation(bot, update.Message.Chat.ID, update.Message.Location)

	// 	return
	// }
}

func New(
	tgBotClient *tgbotapi.BotAPI,
	storage storage.Storage,
	forecast forecast.Forecast, geocode geocode.Geocode,
) *Update {
	return &Update{
		tgBotClient: tgBotClient,

		storage: storage,

		geocode:  geocode,
		forecast: forecast,
	}
}
