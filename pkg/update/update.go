package update

import (
	"strings"
	"twb/pkg/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Update struct {
	storage storage.Storage

	tgBotClient *tgbotapi.BotAPI
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
	default:
	}

	// if update.Message.Text == "now" || update.Message.Text == "for today" ||
	// 	update.Message.Text == "for tomorrow" || update.Message.Text == "for week" ||
	// 	update.Message.Text == "сейчас" || update.Message.Text == "на сегодня" ||
	// 	update.Message.Text == "на завтра" || update.Message.Text == "на неделю" ||
	// 	update.Message.Command() == "now" || update.Message.Command() == "today" ||
	// 	update.Message.Command() == "tomorrow" || update.Message.Command() == "week" {

	// 	WeatherMsgFromCmd(bot, update.Message.Chat.ID, update.Message.Text)

	// 	return
	// }

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

	// WeatherMsgFromCity(bot, update.Message.Chat.ID, update.Message.Text)
}

func New(tgBotClient *tgbotapi.BotAPI, storage storage.Storage) *Update {
	return &Update{
		storage: storage,

		tgBotClient: tgBotClient,
	}
}
