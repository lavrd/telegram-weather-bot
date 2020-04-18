package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"telegram-weather-bot/pkg/bot"
	"telegram-weather-bot/pkg/config"
)

func main() {
	cfg := config.Parse()

	bot, err := bot.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := bot.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-interrupt

	if err := bot.Stop(); err != nil {
		log.Fatal(err)
	}
}
