package bot

import (
	"github.com/pkg/errors"

	"twb/pkg/config"
	"twb/pkg/forecast/openweathermap"
	"twb/pkg/storage"
	"twb/pkg/storage/rethinkdb"
	"twb/pkg/update"

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
	rdb, err := rethinkdb.New(cfg.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize rethinkdb")
	}

	tgBotClient, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize new bot api")
	}

	own := openweathermap.New(cfg.OpenWeatherMapToken)

	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = 60
	updC, err := tgBotClient.GetUpdatesChan(updCfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get telegram update channel")
	}
	upd := update.New(tgBotClient, rdb, own)

	return &Bot{
		storage: rdb,

		update:      upd,
		tgBotClient: tgBotClient,
		updC:        updC,
	}, nil
}
