package errors

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spacelavr/telegram-weather-bot/config"
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
