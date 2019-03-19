package errors

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-weather-bot/pkg/config"
)

func Check(err error) {
	if err != nil {
		if config.Viper.Telegram.Error.Send {
			bot, err := tgbotapi.NewBotAPI(config.Viper.Telegram.Token)
			if err != nil {
				log.Fatal(err)
			}

			msg := tgbotapi.NewMessage(config.Viper.Telegram.Error.Admin, err.Error())

			bot.Send(msg)

			log.Fatal(err)
		}
	}
}
