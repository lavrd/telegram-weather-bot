package bot

import (
	"fmt"
	"telegram-weather-bot/pkg/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	tgBotClient *tgbotapi.BotAPI
	updC        tgbotapi.UpdatesChannel
}

func (b *Bot) Run() error {
	for upd := range b.updC {
		fmt.Println(upd.Message.Text)
	}

	return nil
}

func (b *Bot) Stop() error {
	b.tgBotClient.StopReceivingUpdates()
	b.updC.Clear()

	return nil
}

func New(cfg *config.Config) (*Bot, error) {
	tgBotClient, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}

	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = 60
	updC, err := tgBotClient.GetUpdatesChan(updCfg)
	if err != nil {
		return nil, err
	}

	return &Bot{
		tgBotClient: tgBotClient,
		updC:        updC,
	}, nil
}
