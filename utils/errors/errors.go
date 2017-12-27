package errors

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lavrs/telegram-weather-bot/config"
	"log"
)

// Check check for new errors
func Check(err error) {
	if err != nil {
		// send an error message to a specific user
		bot, _ := tgbotapi.NewBotAPI(config.Cfg.TelegramToken)
		msg := tgbotapi.NewMessage(config.Cfg.MyTelegramID, err.Error())
		bot.Send(msg)
		log.Panic(err)
	}
}
