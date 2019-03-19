package bot

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"telegram-weather-bot/pkg/config"
	"telegram-weather-bot/pkg/db"
	"telegram-weather-bot/pkg/message"
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

	db.Init()

	for update := range updates {
		msg.Updates(bot, update)
	}
}
