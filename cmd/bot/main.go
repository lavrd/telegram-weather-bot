package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"telegram-weather-bot/pkg/bot"
	"telegram-weather-bot/pkg/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.
		New(zerolog.NewConsoleWriter()).
		Level(zerolog.InfoLevel).
		With().Timestamp().
		Logger()

	cfg := config.Parse()

	if cfg.LogLevel != "" {
		logLevel, err := zerolog.ParseLevel(strings.ToLower(cfg.LogLevel))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to parse log level from config")
		}
		log.Logger = log.Logger.Level(logLevel)
	}

	bot, err := bot.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create initialize bot")
	}

	go func() {
		if err := bot.Run(); err != nil {
			log.Fatal().Err(err).Msg("failed to run bot")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-interrupt

	if err := bot.Stop(); err != nil {
		log.Fatal().Err(err).Msg("failed to stop bot")
	}
}
