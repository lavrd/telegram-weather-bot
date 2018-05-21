package bot

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spacelavr/telegram-weather-bot/pkg/config"
	"github.com/spacelavr/telegram-weather-bot/pkg/message"
)

func Daemon() {
	bot, err := tgbotapi.NewBotAPI(config.Viper.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		msg.Updates(bot, update)
	}
}
