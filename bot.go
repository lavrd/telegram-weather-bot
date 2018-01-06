package telegram_weather_bot

import (
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	c "github.com/spacelavr/telegram-weather-bot/config"
	"github.com/spacelavr/telegram-weather-bot/message"
	"github.com/spacelavr/telegram-weather-bot/utils/errors"
)

func main() {
	// configure config
	c.SetConfig()

	// new bot
	bot, err := tgbotapi.NewBotAPI(c.Cfg.TelegramToken)
	errors.Check(err)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://"+c.Cfg.ServerAddr+":8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	// check for updates
	for update := range updates {
		msg.Updates(bot, update)
	}
}
