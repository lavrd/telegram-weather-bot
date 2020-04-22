package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"twb/pkg/bot"
	"twb/pkg/config"
)

func main() {
	log.Logger = zerolog.
		New(zerolog.NewConsoleWriter()).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	if cfg.LogLevel != "" {
		logLevel, err := zerolog.ParseLevel(strings.ToLower(cfg.LogLevel))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to parse log level from config")
		}
		log.Logger = log.Logger.Level(logLevel)
	}

	b, err := bot.New(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create initialize bot")
	}

	go func() {
		if err := b.Run(); err != nil {
			log.Fatal().Err(err).Msg("failed to run bot")
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-interrupt

	if err := b.Stop(); err != nil {
		log.Fatal().Err(err).Msg("failed to stop bot")
	}
}
