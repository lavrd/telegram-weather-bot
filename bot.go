package telegram_weather_bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	c "github.com/lavrs/telegram-weather-bot/config"
	msg "github.com/lavrs/telegram-weather-bot/message"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	"log"
)

func main() {
	// configure config
	c.SetConfig()

	// new bot
	bot, err := tgbotapi.NewBotAPI(c.Cfg.TelegramToken)
	errors.Check(err)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// update every minute
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60
	updates, err := bot.GetUpdatesChan(upd)
	errors.Check(err)

	// check for updates
	for update := range updates {
		msg.Updates(bot, update)
	}
}
