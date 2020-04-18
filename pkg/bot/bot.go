package bot

import (
	"telegram-weather-bot/pkg/config"
	"telegram-weather-bot/pkg/storage"
	"telegram-weather-bot/pkg/storage/rethinkdb"
	"telegram-weather-bot/pkg/update"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

type Bot struct {
	storage storage.Storage

	update      *update.Update
	tgBotClient *tgbotapi.BotAPI
	updC        tgbotapi.UpdatesChannel
}

func (b *Bot) Run() error {
	log.Info().Msg("run telegram weather bot")

	for upd := range b.updC {
		b.update.Handle(&upd)
	}

	return nil
}

func (b *Bot) Stop() error {
	log.Info().Msg("stop telegram weather bot")

	b.tgBotClient.StopReceivingUpdates()
	b.updC.Clear()

	if err := b.storage.Close(); err != nil {
		log.Err(err).Msg("failed to close storage")
	}

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
	update := update.New(tgBotClient)

	storage, err := rethinkdb.New(cfg.DSN)
	if err != nil {
		return nil, err
	}

	return &Bot{
		storage: storage,

		update:      update,
		tgBotClient: tgBotClient,
		updC:        updC,
	}, nil
}
